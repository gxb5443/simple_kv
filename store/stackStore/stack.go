package stackStore

import "errors"

type StackStore struct {
	data map[string]string
}

func (ss *StackStore) Initialize() {
	ss.data = make(map[string]string)
}

func (ss *StackStore) Read(key string) (string, error) {
	if v, ok := ss.data[key]; ok {
		return v, nil
	}
	return errors.New("String not found")
}

func (ss *StackStore) Write(key, val string) error {

}

func (ss *StackStore) Start() error {

}

func (ss *StackStore) Commit() error {

}

func (ss *StackStore) Abort() error {

}
