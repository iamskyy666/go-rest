```go
package main
// demo-1
import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Hello, World!")

	// create a new http client
	client := &http.Client{}
	// res,err:=client.Get("https://jsonplaceholder.typicode.com/posts/1")
	res,err:=client.Get("https://swapi.dev/api/people/2")
	if err!=nil{
		fmt.Println("ERROR making GET request:",err)
		return
	}

	defer res.Body.Close()

	// Read and print the response body
	body,err:=io.ReadAll(res.Body)
	if err!=nil{
		fmt.Println("ERROR :",err)
		return
	}

	fmt.Println(string(body))
}

/*

Hello, World!
{"name":"C-3PO","height":"167","mass":"75","hair_color":"n/a","skin_color":"gold","eye_color":"yellow","birth_year":"112BBY","gender":"n/a","homeworld":"https://swapi.dev/api/planets/1/","films":["https://swapi.dev/api/films/1/","https://swapi.dev/api/films/2/","https://swapi.dev/api/films/3/","https://swapi.dev/api/films/4/","https://swapi.dev/api/films/5/","https://swapi.dev/api/films/6/"],"species":["https://swapi.dev/api/species/2/"],"vehicles":[],"starships":[],"created":"2014-12-10T15:10:51.357000Z","edited":"2014-12-20T21:17:50.309000Z","url":"https://swapi.dev/api/people/2/"}

*/
```