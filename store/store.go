package store

import (
	"errors"
)

type fieldstack []string
type Transactionstack []*Transaction

type StoreEntry struct {
	History fieldstack
}

type StackStore struct {
	data         map[string]StoreEntry
	Transactions Transactionstack
}

type Transaction struct {
	fields map[string]bool
}

func (s fieldstack) Peek() string {
	if len(s) > 0 {
		return s[len(s)-1]
	}
	return ""
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

func (s Transactionstack) Push(v *Transaction) Transactionstack {
	return append(s, v)
}

func (s Transactionstack) Pop() (Transactionstack, *Transaction, error) {
	l := len(s)
	if l == 0 {
		return s, nil, errors.New("Stack Empty")
	}
	return s[:l-1], s[l-1], nil
}

func (ss *StackStore) Initialize() {
	ss.data = make(map[string]StoreEntry)
}

func (ss *StackStore) Read(key string) (string, error) {
	if _, ok := ss.data[key]; ok {
		return ss.data[key].History.Peek(), nil
	}
	return "", errors.New("Key Not Found")
}

func (ss *StackStore) Write(key, val string) {
	if _, ok := ss.data[key]; !ok {
		ss.data[key] = StoreEntry{}
	}
	if len(ss.Transactions) > 0 {
		if ss.Transactions[len(ss.Transactions)-1].fields[key] == false {
			tmpEntry := ss.data[key]
			tmpEntry.History = tmpEntry.History.Push(val)
			ss.data[key] = tmpEntry
			ss.Transactions[len(ss.Transactions)-1].fields[key] = true
		} else {
			ss.data[key].History[len(ss.data[key].History)-1] = val
		}
	} else {
		tmpEntry := ss.data[key]
		tmpEntry.History = tmpEntry.History.Push(val)
		ss.data[key] = tmpEntry
	}
}

func (ss *StackStore) Start() {
	ss.Transactions = ss.Transactions.Push(&Transaction{fields: make(map[string]bool)})
}

func (ss *StackStore) Commit() error {
	var t *Transaction
	if len(ss.Transactions) == 0 {
		return errors.New("No open Transactions")
	}
	ss.Transactions, t, _ = ss.Transactions.Pop()
	for f, _ := range t.fields {
		var tmp string
		tmpEntry := ss.data[f]
		tmpEntry.History, tmp, _ = tmpEntry.History.Pop()
		tmpEntry.History, _, _ = tmpEntry.History.Pop()
		tmpEntry.History = tmpEntry.History.Push(tmp)
		ss.data[f] = tmpEntry
	}
	return nil
}

func (ss *StackStore) Abort() error {
	var t *Transaction
	var err error
	ss.Transactions, t, err = ss.Transactions.Pop()
	if err != nil {
		return errors.New("No open Transactions")
	}
	for f, _ := range t.fields {
		tmpEntry := ss.data[f]
		tmpEntry.History, _, err = tmpEntry.History.Pop()
		ss.data[f] = tmpEntry
		if err != nil {
			delete(ss.data, f)
		}
	}
	return nil
}
