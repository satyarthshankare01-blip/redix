package main

import (
	"fmt"
	"sync"
	"time"
)

type Item struct {
  value string
  expiry time.Time
}

// Store holds the actual data.
// Notice there are NO mutexes here because the channel handles safety.
type Store struct {
	data map[string]Item
  mu   sync.Mutex
}

func NewStore() *Store {
	store := Store{
		data: make(map[string]Item),
	}
  return &store
}

// storing the key in map along with time stamp of when it was stored 
func (s *Store ) set(key string , value string ){
s.data[key] = Item{value , time.Now() }
}



func (s *Store ) get(key string ){
  result , ok  := s.data[key]
  if !ok{
  fmt.Println("the key does not exists")
  return 
  }
  // checking if the key has passed the time limit of expiry if yes then delete it from the map

  if time.Since(result.expiry) > 2 * time.Second{
     delete(s.data , key)
     fmt.Println("your key expired!!")
     return 
  }
  fmt.Println(result)
}


func (s *Store ) delete( key string ){
_ , ok := s.data[key]
if ok {
	delete(s.data , key)
}else{
	fmt.Println("This Key value does not exist")
}
}