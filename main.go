package main

const (
	ErrKeyNotFound  = DictionaryErr("key not found")
	ErrDuplicateKey = DictionaryErr("duplicate key entry")
)

type DictionaryErr string

func (err DictionaryErr) Error() string {
	return string(err)
}

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	value, ok := d[key]
	if !ok {
		return "", ErrKeyNotFound
	}
	return value, nil
}

func (d Dictionary) Add(key, value string) error {
	_, ok := d[key]
	if ok {
		return ErrDuplicateKey
	}
	d[key] = value
	return nil
}

func (d Dictionary) Update(word, definition string) {
	d[word] = definition
}
