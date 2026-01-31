package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"errors"
)

type command  struct {
	args []string

}

// function to handle each of the tcp connection 
func handleConnection (connec net.Conn , ch chan<- command , done <-chan struct{}) {
	defer connec.Close()
	buf := make([]byte , 2024 )

go func( ){
	<-done
	close(ch)
	connec.Close()
}()

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

func checkForData() (*os.File , int) {

file ,err := os.Open("datadb")

	if os.IsNotExist(err){
		fmt.Println("this file does not exist")
		return nil , 1
	}

	if err != nil {
	fmt.Println("Error while opening the file ")
	return nil , 0 
	}

	return file , 1
	
}


func main() {

	sigs := make(chan os.Signal , 1)
	signal.Notify(sigs , os.Interrupt)
	done := make(chan struct{})
    
//Listening for the interrupt signal and closeing the done channel if interrupt is received
	go func () {
     <-sigs
	 close(done)
	}()

    
    dataFile , i := checkForData()

	if dataFile == nil && i == 0 {
	close(sigs)
	close(done)
    return 
	}

	//if dataFile is not nil then we will load data in the map here 


	listener, err := net.Listen("tcp" , ":6379")


	if err != nil {
	fmt.Println("ERROR while listening to the tcp :", err )
	return 

	}

	store := NewStore()


// If the done channel is closed then the server will stop listening to the tcp connection 
	go func(){
    <-done
	listener.Close()
	}()


	for {
	connection , err := listener.Accept()

    if err != nil {
	
	if errors.Is( err , net.ErrClosed ){
    fmt.Println("the listener is closed: SHUTTING DOW THE SERVER")
	return
	}

	
	continue
	}

	fmt.Println("New connection established....")
    
	ch := make( chan command , 20 )

	//starting go routin to read from the command channel to execute desired command 
	go execute(ch , store )

	go handleConnection( connection , ch , done  )
	   
	}

   
}