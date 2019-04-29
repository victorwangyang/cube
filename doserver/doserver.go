package main

import (
	"cube/cluster"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//gClusterInfo is Cluster set information,like map["my-cluster"] = "masterport"
var gClusterInfo map[string]clusterInfo
var gClusterCount int
var gClusterNodesSum int

type clusterInfo struct {
	NodesCount int
	MasterPort string
}

type clusterPutReq struct {
	Clustername string `json:"clustername"`
	Clustersize string `json:"clustersize"`
}

type clusterDeleteReq struct {
	Clustername string `json:"clustername"`
}

func apiV1ClusterPUT(context *gin.Context) {

	var clusterputreq clusterPutReq
	var tempClusterInfo clusterInfo

	err := context.BindJSON(&clusterputreq)
	if err != nil {
		log.Println(err)
		context.JSON(200, gin.H{"errcode": 400, "description": "Post Data Err"})
		return
	}

	tempClusterInfo.MasterPort = cluster.APIV1StartCluster(clusterputreq.Clustersize)
	tempClusterInfo.NodesCount, _ = strconv.Atoi(clusterputreq.Clustersize)
	gClusterInfo[clusterputreq.Clustername] = tempClusterInfo

	gClusterCount = gClusterCount + 1

	tempadded, _ := strconv.Atoi(clusterputreq.Clustersize)
	gClusterNodesSum = gClusterNodesSum + tempadded

	context.JSON(http.StatusOK, gin.H{"clustercount": gClusterCount, "nodescount": gClusterNodesSum})

}

func apiV1ClusterDELETE(context *gin.Context) {

	var clusterdeletereq clusterDeleteReq

	err := context.BindJSON(&clusterdeletereq)
	if err != nil {
		log.Println(err)
		context.JSON(200, gin.H{"errcode": 400, "description": "Post Data Err"})
		return
	}

	cluster.APIV1KillCluster(gClusterInfo[clusterdeletereq.Clustername].MasterPort)

	gClusterCount = gClusterCount - 1

	gClusterNodesSum = gClusterNodesSum - gClusterInfo[clusterdeletereq.Clustername].NodesCount

	delete(gClusterInfo, clusterdeletereq.Clustername)

	context.JSON(http.StatusOK, gin.H{"clustercount": gClusterCount, "nodescount": gClusterNodesSum})

}

func main() {

	gClusterInfo = make(map[string]clusterInfo)
	gClusterCount = 0
	gClusterNodesSum = 0

	route := gin.Default()
	//gin.SetMode(gin.ReleaseMode)
	route.PUT("/api/v1/cluster", apiV1ClusterPUT)
	route.DELETE("/api/v1/cluster", apiV1ClusterDELETE)
	route.Run(":8888")

}
