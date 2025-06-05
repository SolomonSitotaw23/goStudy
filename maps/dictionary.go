package maps

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")

func (d Dictionary) Search(word string) (string, error) {

	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}

// Note on maps
//  -> a nill map (var m map[string]string) will act like an empty map while reading
// but it is a run time error if you tried to write
//  eg m["hello"]="test" -> is a runtime error
// maps should be created like var m = map[string]string{} or var m = make(map[string]string)
