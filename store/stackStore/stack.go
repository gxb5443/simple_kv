package stackStore

import "errors"

type fieldstack []string

type StackStore struct {
	data map[string]fieldstack
}

type transaction struct {
	transaction
	fields []string
}

func (s fieldstack) Push(v string) fieldstack {
	return append(s, v)
}

func (s fieldstack) Pop() (fielstack, string, error) {
	l := len(s)
	if l == 0 {
		return nil, "", errors.New("Stack Empty")
	}
	return s[:l-1], s[l-1], nil
}

func (ss *StackStore) Initialize() {
	ss.data = make(map[string]fieldstack)
}

func (ss *StackStore) Read(key string) (string, error) {
	if v, ok := ss.data[key]; ok {
		ss.data[key], out, e = ss.data[key].Pop()
		if e != nil {
			return "", e
		}
		return out, nil
	}
	return "", errors.New("String not found")
}

func (ss *StackStore) Write(key, val string) {
	ss.data[key] = val
}

func (ss *StackStore) Start() error {

}

func (ss *StackStore) Commit() error {

}

func (ss *StackStore) Abort() error {

}
