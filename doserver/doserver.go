package main

import (
	"cube/cluster"
	"cube/restapi"
	"log"
	"net/http"
)

//ClusterInfo is Cluster set information
type ClusterInfo struct {
	ClusterName string
	NodeNumber  string
}

var webServer = []restapi.RestfulInterface{
	{Path: "/api/v1/cluster", Handle: apiV1Cluster},
}

func apiV1Cluster(w http.ResponseWriter, r *http.Request) {

	if r.Method == "PUT" {

		cluster.APIV1StartCluster("my-cluster", "3")
		log.Println(cluster.GClusterStatus)

	} else if r.Method == "DELETE" {
		cluster.APIV1KillCluster("dddd")

	}
	// var clusterinfo = ClusterInfo{}

	// body, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// log.Println("body:", string(body))
	// json.Unmarshal(body, &clusterinfo)
	// //log.Println("body:", body)

}

func main() {

	restapi.InitRestSvr("8888", webServer)

}
