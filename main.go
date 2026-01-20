package main

import (
	"fmt"
	"net"
	"strconv"
)

type command  struct {
	args []string

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
	
	// parsing logic for extracting the command from request 
	var comm command 
	i := 0 
	for range n-1 {	

	if string(buf[i]) == "$"{

    nb ,_ := strconv.Atoi(string(buf[i+1]))

    fmt.Println(string(buf[i+4:i+4+nb]))
	comm.args = append( comm.args , string(buf[i+4:i+4+nb]) )

   }

 i++      
}
connec.Write([]byte("+OK\r\n"))

//passing the command into command chaannel 
ch<-comm

}

}

func main() {
	listener, err := net.Listen("tcp" , ":6379")
	if err != nil {
		fmt.Println("ERROR while listening to the tcp :", err )
		return 
	}

	store := NewStore()


	for{

	connection , err := listener.Accept()
    if err != nil {
	fmt.Println("Failed to Establish the Connection: ", err )
	continue
	} 
	fmt.Println("New connection established....")
    
	ch := make( chan command , 20 )

	//starting go routin to read from the command channel to execute desired command 
	go execute(ch , *store  )

	go handleConnection( connection , ch  )
	   
	}
    
}