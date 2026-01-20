

package main

// Store holds the actual data.
// Notice there are NO mutexes here because the channel handles safety.
type Store struct {
	data map[string]string
}

func NewStore() *Store {
	return &Store{
		data: make(map[string]string),
	}
}