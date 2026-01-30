package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

type command  struct {
	args []string

}

// function to handle each of the tcp connection 
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

//passing the command into command channel 
ch<-comm

}

}

// function to look if there is pre-existing data is present for the redis in dedicated file or is the first time boot 
// of this redis application

func checkForData() *os.File {

file ,err := os.Open("datadb")

	if os.IsNotExist(err){
		fmt.Println("this file does not exist")
		return nil
	}

	if err != nil {
		fmt.Println("Error while opening the file ")
		
	}

	return file
}


func main() {
    
    //dataFile := checkForData()
    
	//if dataFile is not nil then we will load data in the map here 

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
	go execute(ch , store )

	go handleConnection( connection , ch  )
	   
	}


    
}