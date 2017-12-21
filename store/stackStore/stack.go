package stackStore

import "errors"

type fieldstack []string
type transactionstack []Transaction

type StoreEntry struct {
	History fieldstack
}

type StackStore struct {
	data         map[string]StoreEntry
	transactions transactionstack
}

type Transaction struct {
	fields map[string]bool
}

func (s fieldstack) Peek() string {
	return s[len(s)-1]
}
func (s fieldstack) Push(v string) fieldstack {
	return append(s, v)
}

func (s fieldstack) Pop() (fieldstack, string, error) {
	l := len(s)
	if l == 0 {
		return nil, "", errors.New("Stack Empty")
	}
	return s[:l-1], s[l-1], nil
}

func (s transactionstack) Push(v string) transactionstack {
	return append(s, v)
}

func (s transactionstack) Pop() (transactionstack, string, error) {
	l := len(s)
	if l == 0 {
		return nil, "", errors.New("Stack Empty")
	}
	return s[:l-1], s[l-1], nil
}

func (ss *StackStore) Initialize() {
	ss.data = make(map[string]StoreEntry)
}

func (ss *StackStore) Read(key string) (string, error) {
	if v, ok := ss.data[key]; ok {
		return ss.data[key].History.Peek(), nil
	}
	return "", errors.New("Key Not Found")
}

func (ss *StackStore) Write(key, val string) {
	if v, ok := ss.data[key]; !ok {
		ss.data[key] = new(StoreEntry)
	}
	if len(ss.transactions) > 0 {
		if ss.transactions[len(ss.transactions)-1][key] == false {
			ss.data[key].History.Push(val)
			ss.transactions[len(ss.transactions)-1][key] = true
		} else {
			ss.data.History[key][len(ss.data[key].History-1)] = val
		}
	} else {
		ss.data.History[key][len(ss.data[key].History-1)] = val
	}
}

func (ss *StackStore) Start() {
	ss.transactions.Push(new(Transaction))
}

func (ss *StackStore) Commit() error {
	fields_to_commit := ss.transactions.Pop()
	for f := range fields_to_commit {
		ss.data[f].History, tmp, _ = ss.data[f].History.Pop()
		ss.data[f].History, _, _ = ss.data[f].History.Pop()
		ss.data[f].Push(tmp)
	}
}

func (ss *StackStore) Abort() error {
	ss.transactions, fields_to_abort, err = ss.transactions.Pop()
	if err != nil {
		return errors.New("No open Transactions")
	}
	for f := range fields_to_commit {
		ss.data[f].History, _, err = ss.data[f].History.Pop()
		if err != nil {
			delete(ss.data, f)
		}
	}
}
