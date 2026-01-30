package main

import "os"

func snapShot(store *Store) map[string]Item {
	store.mu.Lock()
	defer store.mu.Unlock()
	clone := make(map[string]Item)
	for key , value := range store.data{
		clone[key] = value
	}

	return clone 
}

