```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

// Demo-2
// create server, receive req.,send resp.

func main() {

	http.HandleFunc("/",func (resp http.ResponseWriter, req *http.Request){
		fmt.Fprintln(resp,"Hello Server ✅")
	})
	
	// PORT:="127.0.0.1:3000" // localhost
	 PORT:=":3000" // better

	fmt.Println("✅ Server is listening on PORT:",PORT)
	err:=http.ListenAndServe(PORT,nil)
	if err!=nil{
		log.Fatalln("⚠️ERR:",err)
	}
}

// $ go run .
// ✅ Server is listening on PORT: :3000
```