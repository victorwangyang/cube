package restapi

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

//RestfulInterface is to define the process api request
type  RestfulInterface struct {
	Path string
	Handle func(http.ResponseWriter,*http.Request)
}

//InitRestSvr is to prepare resource of path/func and start to listen at the addr
func InitRestSvr(port string,register []RestfulInterface) {

	r := mux.NewRouter()

	// Routes consist of a path and a handler function.
	
	
	for index := range register {
		
		r.HandleFunc(register[index].Path, register[index].Handle)
	}

	http.Handle("/", r)

	log.Println(" Listening.....", port)
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":" + port, r))
	
}
