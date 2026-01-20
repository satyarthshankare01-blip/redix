package main 
import (
	"fmt"
)
func (s *Store ) set(key string , value string ){
s.data[key] = value
}


func (s *Store ) get(key string   ){
  result , ok  := s.data[key]
  if ok{
  fmt.Println(result)
  }
}


func (s *Store ) delete( key string ){
_ , ok := s.data[key]
if ok {
	delete(s.data , key)
}else{
	fmt.Println("This Key value does not exist")
}
}