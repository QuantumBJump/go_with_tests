package dictionary

type Dictionary map[string]string

const (
	ErrWordNotFound       = DictionaryError("could not find the word you were looking for")
	ErrWordExists         = DictionaryError("this word already exists")
	ErrUpdateWordNotExist = DictionaryError("failed to update word because it does not exist")
)

type DictionaryError string

func (e DictionaryError) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	def, ok := d[word]

	if !ok {
		return "", ErrWordNotFound
	}
	return def, nil
}

func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case ErrWordNotFound:
		d[word] = def // ErrNotFound means we can add the word
	case nil:
		return ErrWordExists // No err from search means the word is already in the dictionary
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case ErrWordNotFound:
		return ErrUpdateWordNotExist
	case nil:
		d[word] = def
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
