package cluster

import (
	"log"
	"math/rand"
	"net/rpc"
	"os"
	"os/exec"
	"strconv"
)

// constant of cluster
const (
	TargetAddr      = "http://localhost"
	MasterDirectory = "/api/v1/cluster/master"
	NodeDirectory   = "/api/v1/cluster/node"

	MasterExePosition = "../cluster/master/master"
	NodeExePosition   = "../cluster/node/node"
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

//GNodeInfo is GNodeInfo
var GNodeInfo map[string]NodeInfo

// APIV1StartCluster is called by doserver's RESTful API to start the Cluster
func APIV1StartCluster(nodenumber string) string {

	//set cluster staus values
	mastername := "Master-PM-" + strconv.Itoa(80000+rand.Intn(10000))
	masterport := strconv.Itoa(30000 + rand.Intn(10000))

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

		nodename := "Node-PM-" + strconv.Itoa(90000+rand.Intn(10000))
		nodeport := strconv.Itoa(20000 + rand.Intn(10000))

		var tempNodeInfo = NodeInfo{nodeport, 0}
		GNodeInfo[nodename] = tempNodeInfo

		cmd := exec.Command(NodeExePosition, nodename, masterport, nodeport)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Start()

	}

	return
}
