package cluster

import (
	"log"
	"math/rand"
	"net/rpc"
	"os"
	"os/exec"
	"strconv"
	"sync"
)

// constant of cluster
const (
	TargetAddr      = "http://localhost"
	MasterDirectory = "/api/v1/cluster/master"
	NodeDirectory   = "/api/v1/cluster/node"

	MasterExePosition = "./bin/master"
	NodeExePosition   = "./bin/node"
)

// MasterInfo is struct of Master information
type MasterInfo struct {
	MasterName string
	MasterPort string
	NodeNumber string
}

// NodeInfo is struct of Node information
type NodeInfo struct {
	NodePort      string
	NodeLiveCount int
}

//GMasterInfo is GMasterInfo
var GMasterInfo MasterInfo

//GNodeLiveCount is GNodeLiveCount
var GNodeLiveCount sync.Map

//GNodePort is GNodePort
var GNodePort map[string]string

// APIV1StartCluster is called by doserver's RESTful API to start the Cluster
func APIV1StartCluster(nodenumber string) string {

	//set cluster staus values
	masterport := strconv.Itoa(30000 + rand.Intn(10000))
	mastername := "Master-PM-" + masterport

	//start Master process
	cmd := exec.Command(MasterExePosition, mastername, masterport, nodenumber)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Start()

	return masterport
}

// APIV1KillCluster is called by doserver's RESTful API to kill the Cluster
func APIV1KillCluster(masterport string) {

	var live, reply bool

	live = false

	//kill master,before that ,master will kill nodes firstly
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:"+masterport)

	if err != nil {
		log.Fatal("dialing:", err)
	}

	defer client.Close()

	err = client.Call("Master.KillMaster", live, &reply)

	if err != nil {
		log.Fatal("arith error:", err)
	}

	return
}

//APIV1StartNodesDeamon is called by master's main() to start noedes
func APIV1StartNodesDeamon(nodenumber string, masterport string) {

	nodeNumber, _ := strconv.Atoi(nodenumber)

	for i := 0; i < nodeNumber; i++ {

		nodeport := strconv.Itoa(20000 + rand.Intn(10000))
		nodename := "Node-PM-" + nodeport

		GNodePort[nodename] = nodeport
		GNodeLiveCount.Store(nodename, 0)

		cmd := exec.Command(NodeExePosition, nodename, masterport, nodeport)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Start()

	}

	return
}
