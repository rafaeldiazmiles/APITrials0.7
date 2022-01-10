package main

import (
	jsonparse "encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)

//  Args holds arguments passed to json rpc service
type Args struct {
	Id string
}

// Book struct holds Book JSON structure
type Book struct {
	Id     string `json:"string,omitempty"`
	Name   string `json:"name,omitempty"`
	Author string `json:"author,omitempty"`
}

type JSONServer struct{}

// GiveBookDetail
func (t *JSONServer) GiveBookDetail(r *http.Request, args *Args, reply *Book) error {
	var books []Book
	// Read JSON file and load data
	raw, readErr := ioutil.ReadFile("./books.json")
	if readErr != nil {
		fmt.Println("Hola rafiii")
		log.Println("error:", readErr)
		os.Exit(1)
	}
	// Unmarshal JSON raw data into books array
	marshalErr := jsonparse.Unmarshal(raw, &books)
	if marshalErr != nil {
		log.Println("error:", marshalErr)
		os.Exit(2)
	}
	// itera los libros para encontrar el que buscas
	for _, book := range books {
		if book.Id == args.Id {
			// encontraste el libro capo
			*reply = book
			break
		}
	}
	return nil
}

func main() {
	//  Asi que aca no se que de crear el nuevo rpc server
	s := rpc.NewServer() // no se que de registrar la data como JSOn
	s.RegisterCodec(json.NewCodec(), "application/json")
	// Register the service by creating a new JSON server
	s.RegisterService(new(JSONServer), "")
	r := mux.NewRouter()
	r.Handle("/rpc", s)
	http.ListenAndServe(":1234", r)

}
