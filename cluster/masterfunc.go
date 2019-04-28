package cluster

import (
	"log"
	"net/rpc"
	"os"
	"time"
)

//GMasterlive is the flag of liveness
var GMasterlive = true

//Master is Master
type Master int

//HeartBeatInfo is NodeHeartBeatInfo
type HeartBeatInfo struct {
	NodeName      string
	NodeLiveCount int
}

//HeartBeatNotify is func for regester
func (m *Master) HeartBeatNotify(Heartbeatinfo *HeartBeatInfo, reply *bool) error {

	// nodename := (*Heartbeatinfo).NodeName

	// var tempNodeInfo = NodeInfo{GNodeInfo[nodename].NodePort, (*Heartbeatinfo).NodeLiveCount}

	// GNodeInfo[nodename] = tempNodeInfo

	log.Printf("beat heart..........")

	*reply = true

	return nil
}

//KillMaster is KillMaster
func (m *Master) KillMaster(masterlive *bool, reply *bool) error {

	//kill nodes
	for _, v := range GNodeInfo {

		client, err := rpc.DialHTTP("tcp", "127.0.0.1:"+v.NodePort)
		if err != nil {
			log.Fatal("dialing:", err)
		}

		var live, reply bool

		live = *masterlive

		err = client.Call("Node.KillNode", live, &reply)

		if err != nil {
			log.Fatal("arith error:", err)
		}

	}

	// kill master
	GMasterlive = *masterlive

	*reply = true

	return nil
}

// MasterExit is Master Process exit function
func MasterExit() {

	for {
		if GMasterlive == true {
			time.Sleep(time.Second * 2)
		} else {
			log.Printf("%s stopping ......", GMasterInfo.MasterName)
			os.Exit(0)
		}
	}
}

/*
//MasterDeleteNode is to Process request from Client
func MasterDeleteNode(r *http.Request) {


	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}


	nodesNumber, _ := strconv.Atoi(GCluster.NodesNumber)


	for i := 0; i < nodesNumber; i++ {
		if GCluster.Nodes[i] == string(body){

		}
		port := strconv.Itoa(InitNodePort + i)
		StartNode(i, conf.MachineNamelist[i], port)
	}

}
*/
