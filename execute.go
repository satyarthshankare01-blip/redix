package main
import (
	"strings"
)

func execute( ch <-chan command , store *Store ) {

	 for cm := range ch {
      
	switch strings.ToUpper(cm.args[0]){
	case "SET":
		store.set(cm.args[1] , cm.args[2]  )

	case "GET":
		store.get(cm.args[1]   )

	case "DELETE":
         store.delete(cm.args[1]   )

	}
	

	 }
}