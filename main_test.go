package main

import (
	"testing"
)

func assertString(t testing.TB, want, got string) {
	t.Helper()
	if got != want {
		t.Errorf("want %q, got %q", want, got)
	}
}
func assertError(t testing.TB, want, got error) {
	t.Helper()
	if got != want {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestSearch(t *testing.T) {
	t.Run("known word", func(t *testing.T) {
		myMap := Dictionary{"test": "this is just a test"}
		got, _ := myMap.Search("test")
		want := "this is just a test"

		assertString(t, want, got)
	})
	t.Run("unknown word", func(t *testing.T) {
		myMap := Dictionary{"test": "this is just a test"}
		_, err := myMap.Search("unknown")

		assertError(t, ErrKeyNotFound, err)
	})
}

func assertDefinition(t testing.TB, m Dictionary, key, value string) {
	t.Helper()
	got, err := m.Search(key)

	if err != nil {
		t.Fatal("should find added word: ", key)
	}

	if got != value {
		t.Errorf("want %q, got %q, given %q", value, got, key)
	}
}

func TestAdd(t *testing.T) {
	t.Run("add new", func(t *testing.T) {
		myMap := Dictionary{"test": "this is just a test"}
		key, value := "cat", "this is a cat"
		err := myMap.Add(key, value)
		assertError(t, nil, err)
		assertDefinition(t, myMap, key, value)
	})

	t.Run("add existing", func(t *testing.T) {
		key, value := "test", "this is just a test"
		myMap := Dictionary{"test": value}
		err := myMap.Add(key, "this is a cat")
		assertError(t, ErrDuplicateKey, err)
		assertDefinition(t, myMap, key, value)
	})
}

func TestUpdate(t *testing.T) {
	word := "test"
	definition := "this is just a test"
	dictionary := Dictionary{word: definition}
	newDefinition := "new definition"

	dictionary.Update(word, newDefinition)

	assertDefinition(t, dictionary, word, newDefinition)
}
