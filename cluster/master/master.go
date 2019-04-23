package main

import (
	"cube/cluster"
	"cube/restapi"
	"log"
	"net/http"
	"os"
	//	"strconv"
)

var masterAPI = []restapi.RestfulInterface{
	{Path: cluster.MasterDirectory, Handle: apiV1MasterHandler},
}

func apiV1MasterHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		log.Println("GET.....from Master")
	case "POST":
		{
			log.Printf("master exit......")
			cluster.GMasterlive = false
		}
	default:
		log.Println("DEFAULT.......")

	}
}

func main() {

	go cluster.MasterExit()

	restapi.InitRestSvr(os.Args[2], masterAPI)
}
