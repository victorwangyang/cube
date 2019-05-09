package cluster

import (
	"log"
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

var gNodesNotify = true

//HeartBeatNotify is func for regester
func (m *Master) HeartBeatNotify(Heartbeatinfo *HeartBeatInfo, reply *bool) error {

	GNodeLiveCount.Store((*Heartbeatinfo).NodeName, (*Heartbeatinfo).NodeLiveCount)

	//	log.Printf(" %s beat heart..........", (*Heartbeatinfo).NodeName)

	*reply = gNodesNotify

	return nil
}

//KillMaster is KillMaster
func (m *Master) KillMaster(masterlive *bool, reply *bool) error {

	gNodesNotify = *masterlive
	time.Sleep(time.Second * 6)

	GMasterlive = *masterlive

	*reply = true

	return nil
}

// MasterExit is Master Process exit function
func MasterExit() {

	for {
		if GMasterlive == true {
			time.Sleep(time.Second * 5)
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
