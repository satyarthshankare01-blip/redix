package main

import (
	"fmt"
	"net"
	"strconv"
)

type command  struct {
	cm string
	arg []string

}

func handleConnection (connec net.Conn , ch chan<- command ) {
	defer connec.Close()
	buf := make([]byte , 2024 )

    for {
	n , err := connec.Read(buf)
		
	if err != nil {
	fmt.Println("closing the connection: " , err )
	break
	}
	
	i := 0 
	for  range n  {
	fmt.Println(string(buf[i]))
	i++
	}
	
	if string(buf[0]) == "*"{
    val , _ := strconv.Atoi(string(buf[1]))
	
	j := 1+4
	for range val {
	if()
    nbtR,_ := strconv.Atoi(string(buf[i+3]))

	str := string(bug[])
	}

	}
	


	connec.Write([]byte("+OK\r\n"))

        
	 }

}

func main() {
	listener, err := net.Listen("tcp" , ":6379")
	if err != nil {
		fmt.Println("ERROR while listening to the tcp :", err )
		return 
	}

	for{

	connection , err := listener.Accept()
    if err != nil {
	fmt.Println("Failed to Establish the Connection: ", err )
	continue
	} 
	fmt.Println("New connection established....")
    
	ch := make( chan command , 20 )
	go handleConnection( connection , ch  )
	   
	}
    

}