package main

import (
	"cube/cluster"
	"log"

	"github.com/gin-gonic/gin"
)

//gClusterInfo is Cluster set information,like map["my-cluster"] = "masterport"
var gClusterInfo map[string]string

type clusterPutReq struct {
	Clustername string `json:"clustername"`
	Clustersize string `json:"clustersize"`
}

func apiV1ClusterPUT(context *gin.Context) {

	var clusterputreq clusterPutReq

	err := context.BindJSON(&clusterputreq)
	if err != nil {
		log.Println(err)
		context.JSON(200, gin.H{"errcode": 400, "description": "Post Data Err"})
		return
	}

	gClusterInfo[clusterputreq.Clustername] = cluster.APIV1StartCluster(clusterputreq.Clustersize)

	log.Println(gClusterInfo[clusterputreq.Clustersize])
	//log.Println(gClusterInfo[clustername])

}

func apiV1ClusterDELETE(context *gin.Context) {

	clustername := "my-cluster"

	cluster.APIV1KillCluster(gClusterInfo[clustername])

}

func main() {

	gClusterInfo = make(map[string]string)
	route := gin.Default()
	//gin.SetMode(gin.ReleaseMode)
	route.PUT("/api/v1/cluster", apiV1ClusterPUT)
	route.DELETE("/api/v1/cluster", apiV1ClusterDELETE)
	route.Run(":8888")

}
