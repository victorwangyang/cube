package cluster

import (
	"cube/restapi"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
)

// constant of cluster
const (
	TargetAddr      = "http://localhost"
	MasterDirectory = "/api/v1/cluster/master"
	NodeDirectory   = "/api/v1/cluster/node"

	MasterExePosition = "./cluster/master/master"
	NodeExePosition   = "./cluster/node/node"
)

//GClusterStatus is saving the info of cluster
var GClusterStatus = V1ClusterStatus{}

// V1ClusterStatus is struct of cluster status
type V1ClusterStatus struct {
	ClusterName string
	MasterName  string
	MasterPort  string
	NodeNumber  string
	Nodes       []V1NodeStatus
}

// V1NodeStatus is struct of cluster status file
type V1NodeStatus struct {
	NodePort string
	NodeName string
}

// APIV1StartCluster is to start a cluster of V1
func APIV1StartCluster(clustername string, nodenumber string) {

	//set cluster staus values
	GClusterStatus.ClusterName = clustername

	randNum := strconv.Itoa(rand.Intn(100000))
	GClusterStatus.MasterName = "Master-PM-" + randNum

	GClusterStatus.MasterPort = strconv.Itoa(30000 + rand.Intn(10000))

	GClusterStatus.NodeNumber = nodenumber

	nodenum, _ := strconv.Atoi(nodenumber)
	GClusterStatus.Nodes = make([]V1NodeStatus, nodenum, nodenum)

	//start Master process
	cmd := exec.Command(MasterExePosition, GClusterStatus.MasterName, GClusterStatus.MasterPort)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Start()

	log.Printf("%s is starting......", GClusterStatus.MasterName)

	//Start Nodes process
	v1startNodesDeamon(&GClusterStatus)

	return
}

// APIV1KillCluster is to kill the Cluster
func APIV1KillCluster(clustername string) {

	var body []byte
	nodeNumber, _ := strconv.Atoi(GClusterStatus.NodeNumber)

	// kill nodes first
	for i := 0; i < nodeNumber; i++ {
		log.Printf("%s  is Stopping......", GClusterStatus.Nodes[i].NodeName)
		restapi.Post(TargetAddr+":"+GClusterStatus.Nodes[i].NodePort+NodeDirectory, body)
	}

	//then kill master last
	restapi.Post(TargetAddr+":"+GClusterStatus.MasterPort+MasterDirectory, body)
	log.Printf("%s is stopping......", GClusterStatus.MasterName)
	return
}

func v1startNodesDeamon(status *V1ClusterStatus) {

	nodeNumber, _ := strconv.Atoi(status.NodeNumber)

	for i := 0; i < nodeNumber; i++ {

		randNum := strconv.Itoa(rand.Intn(100000))
		status.Nodes[i].NodeName = "Node-PM-" + randNum
		status.Nodes[i].NodePort = strconv.Itoa(50000 + rand.Intn(10000))

		cmd := exec.Command(NodeExePosition, status.Nodes[i].NodeName, status.Nodes[i].NodePort, status.MasterPort)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Start()

		log.Printf("%s is Starting......", status.Nodes[i].NodeName)
	}

	return
}
