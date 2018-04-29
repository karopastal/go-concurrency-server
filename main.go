package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// MyResp is awesome
type MyResp struct {
	Name  string `json:"name"`
	State bool   `json:"state"`
}

// HelloServer serving hello
func HelloServer(res http.ResponseWriter, req *http.Request) {
	name := mux.Vars(req)["name"]

	respObj := MyResp{
		Name:  name,
		State: true,
	}

	respJSON, err := json.Marshal(respObj)

	if err != nil {
		panic(err)
	}

	res.Header().Set("Content-type", "application/json")
	res.Write(respJSON)
}

func main() {
	r := mux.NewRouter()

	r.Path("/{name}").Methods("GET").HandlerFunc(HelloServer)

	n := negroni.Classic()

	n.UseHandler(r)
	n.Run(":3030")
}
