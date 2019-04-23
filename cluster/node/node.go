package main

import (
	"cube/cluster"
	"cube/restapi"
	"log"
	"net/http"
	"os"
)

func apiV1NodeHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		log.Println("GET.....from Master")
	case "POST":
		{
			log.Printf("node exit......")
			cluster.GNodelive = false

		}
	default:
		log.Println("DEFAULT.......")

	}

}

var nodeAPI = []restapi.RestfulInterface{
	{Path: cluster.NodeDirectory, Handle: apiV1NodeHandler},
}

func main() {

	//var NodeName = os.Args[1]
	var PortNum = os.Args[2]

	go cluster.NodeExit()
	restapi.InitRestSvr(PortNum, nodeAPI)

}
