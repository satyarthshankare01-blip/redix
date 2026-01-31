package main

import (
	"encoding/binary"
	"os"
)

func snapShot(store *Store) map[string]Item {
	store.mu.Lock()
	defer store.mu.Unlock()
	clone := make(map[string]Item , len(store.data) )
	for key , value := range store.data{
		clone[key] = value
	}

	return clone 
}



func saveSnapshot(store *Store) {

	copy := snapShot(store)
	filename := "tempdata"


	file, err  := os.Create(filename)
	if err != nil {
	return 
	}
	defer file.Close()
    
	for key , item := range copy {

		keyBytes := []byte(key)
		binary.Write(file , binary.LittleEndian , uint32(len(keyBytes)))
		file.Write(keyBytes)

		valueBytes := []byte(item.value)
		binary.Write(file , binary.LittleEndian , uint32(len(valueBytes)))
		file.Write(valueBytes)
	}

	file.Write([]byte{0xFF})

	os.Rename(filename , "datadb")


    

}