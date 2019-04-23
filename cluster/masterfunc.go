package cluster

import (
	"os"
	"time"
)

//GMasterlive is the flag of liveness
var GMasterlive = true

// MasterExit is Master Process exit function
func MasterExit() {

	for {
		if GMasterlive == true {
			time.Sleep(time.Second * 2)
		} else {
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
