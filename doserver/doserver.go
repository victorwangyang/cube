package main

import (
	"cube/cluster"
	"log"

	"github.com/gin-gonic/gin"
)

//gClusterInfo is Cluster set information,like map["my-cluster"] = "masterport"
var gClusterInfo map[string]string

func apiV1ClusterPUT(context *gin.Context) {

	clustername := "my-cluster"
	clustersize := "3"
	gClusterInfo[clustername] = cluster.APIV1StartCluster(clustersize)

	log.Println(gClusterInfo[clustername])

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
