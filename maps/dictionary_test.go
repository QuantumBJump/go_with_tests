package dictionary

import "testing"

func TestDictionaryError(t *testing.T) {
	err := DictionaryError("test")
	got := err.Error()
	want := "test"

	assertStrings(t, got, want)
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, err := dictionary.Search("test")
		want := "this is just a test"

		assertNoError(t, err)
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("notaword")

		assertError(t, err, ErrWordNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dict := Dictionary{}
		word := "test"
		def := "this is just a test"

		dict.Add(word, def)

		assertDefinition(t, dict, word, def)
	})
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		def := "this is just a test"
		dict := Dictionary{word: def}
		anotherDef := "another definition!"
		err := dict.Add(word, anotherDef)

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dict, word, def)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		def := "this is just a test"
		dict := Dictionary{word: def}
		newDef := "new definition"

		err := dict.Update(word, newDef)

		assertError(t, err, nil)
		assertDefinition(t, dict, word, newDef)
	})
	t.Run("word does not exist", func(t *testing.T) {
		dict := Dictionary{}
		word := "test"
		newDef := "new definition"

		err := dict.Update(word, newDef)

		assertError(t, err, ErrUpdateWordNotExist) // assert Update returns an error
		got, err := dict.Search(word)
		assertError(t, err, ErrWordNotFound) // assert Update does not create the word
		assertStrings(t, got, "")
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dict := Dictionary{word: "test definition"}

	dict.Delete(word)

	_, err := dict.Search(word)
	if err != ErrWordNotFound {
		t.Errorf("Expected word to be deleted: %q", word)
	}
}

func assertDefinition(t testing.TB, dict Dictionary, word, def string) {
	t.Helper()

	got, err := dict.Search(word)
	if err != nil {
		t.Fatalf("unexpected error: %q", err)
	}

	if got != def {
		t.Errorf("got %q, want %q", got, def)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNoError(t testing.TB, err error) {
	if err != nil {
		t.Errorf("unexpected error: %q", err.Error())
	}
}
