package cluster

import (
	"log"
	"net/rpc"
	"os"
	"time"
)

//Node is Node
type Node int

//GNodelive is the flag of liveness
var GNodelive = true

//KillNode is func for regester
func (n *Node) KillNode(nodelive *bool, reply *bool) error {

	GNodelive = *nodelive
	*reply = true

	return nil
}

// NodeExit is node keep live function
func NodeExit() {

	for {
		if GNodelive == true {
			time.Sleep(time.Second * 2)
		} else {
			os.Exit(0)
		}

	}
}

//NodeLiveHeartBeat is NodeLiveHeartBeat
func NodeLiveHeartBeat(masterport string, nodename string) {

	var reply bool
	var heartbeatinfo HeartBeatInfo

	heartbeatinfo.NodeName = nodename
	heartbeatinfo.NodeLiveCount = 5

	for {
		client, err := rpc.DialHTTP("tcp", "127.0.0.1:"+masterport)
		if err != nil {
			log.Fatal("dialing:", err)
		}

		err = client.Call("Master.HeartBeatNotify", heartbeatinfo, &reply)

		if err != nil {
			log.Fatal("arith error:", err)
		}

		time.Sleep(time.Second * 5)

	}

}
