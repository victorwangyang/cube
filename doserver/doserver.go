package main

import (
	"cube/cluster"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

//gClusterInfo is Cluster set information,like map["my-cluster"] = "masterport"
var gClusterInfo map[string]clusterInfo
var gClusterCount int
var gClusterNodesSum int

type clusterInfo struct {
	NodesCount int
	MasterPort string
	CreateTime string
}

type clusterPutReq struct {
	Clustername string `json:"clustername"`
	Clustersize string `json:"clustersize"`
}

type clusterDeleteReq struct {
	Clustername string `json:"clustername"`
}

type clustergetres struct {
	Clustername string `json:"clustername"`
	Masterport  int    `json:"masterport"`
	Nodescount  int    `json:"nodescount"`
	Createtime  string `json:"createtime"`
}

type statgetres struct {
	Activecluster int `json:"activecluster"`
	Activenodes   int `json:"activenodes"`
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
	tempClusterInfo.CreateTime = time.Now().Format("2006-01-02 15:04")

	log.Println(tempClusterInfo.MasterPort)
	log.Println(clusterputreq.Clustername)

	gClusterInfo[clusterputreq.Clustername] = tempClusterInfo

	gClusterCount = gClusterCount + 1

	tempadded, _ := strconv.Atoi(clusterputreq.Clustersize)
	gClusterNodesSum = gClusterNodesSum + tempadded

	context.JSON(http.StatusOK, gin.H{"clustercount": gClusterCount, "nodescount": gClusterNodesSum})

}

func apiV1ClusterDELETE(context *gin.Context) {

	clustername := context.Param("name")

	// log.Println(clustername)
	// log.Println(gClusterInfo[clustername].MasterPort)

	cluster.APIV1KillCluster(gClusterInfo[clustername].MasterPort)

	gClusterCount = gClusterCount - 1

	gClusterNodesSum = gClusterNodesSum - gClusterInfo[clustername].NodesCount

	delete(gClusterInfo, clustername)

	context.JSON(http.StatusOK, gin.H{"clustercount": gClusterCount, "nodescount": gClusterNodesSum})

}

func apiV1ClusterGET(context *gin.Context) {

	var resmsg []clustergetres = make([]clustergetres, gClusterCount, gClusterCount)
	var i = 0

	for k, v := range gClusterInfo {
		resmsg[i].Clustername = k
		resmsg[i].Masterport, _ = strconv.Atoi(v.MasterPort)
		resmsg[i].Nodescount = v.NodesCount
		resmsg[i].Createtime = v.CreateTime
		i++
	}

	context.JSON(http.StatusOK, resmsg)
}

func apiV1StatisticGET(context *gin.Context) {

	var stat statgetres

	stat.Activecluster = gClusterCount
	stat.Activenodes = gClusterNodesSum

	context.JSON(http.StatusOK, stat)
}

func main() {

	gClusterInfo = make(map[string]clusterInfo)
	gClusterCount = 0
	gClusterNodesSum = 0

	route := gin.Default()
	//gin.SetMode(gin.ReleaseMode)
	route.PUT("/api/v1/cluster", apiV1ClusterPUT)
	route.DELETE("/api/v1/cluster/:name", apiV1ClusterDELETE)
	route.GET("/api/v1/cluster", apiV1ClusterGET)
	route.GET("api/v1/cluster/statistic", apiV1StatisticGET)
	route.Run(":8888")

}
