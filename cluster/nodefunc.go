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

//GMasterPort is the flag of liveness
var GMasterPort string

//GNodeName is the flag of liveness
var GNodeName string

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
			log.Printf("%s stopping ......", GNodeName)
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

	time.Sleep(time.Second * 3)

	client, err := rpc.DialHTTP("tcp", "127.0.0.1:"+masterport)
	if err != nil {
		log.Fatal("NodeLiveHeartBeat:", err)
	}

	defer client.Close()

	for {

		err = client.Call("Master.HeartBeatNotify", heartbeatinfo, &reply)

		if reply == false {

			GNodelive = false
		}

		if err != nil {
			log.Println("NodeLiveHeartBeat error:", err)
		}

		time.Sleep(time.Second * 5)

	}

}
