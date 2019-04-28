package main

import (
	"cube/cluster"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
)

func main() {

	cluster.GNodeName = os.Args[1]
	cluster.GMasterPort = os.Args[2]
	nodeport := os.Args[3]

	//start a thread to watch live state of Node
	go cluster.NodeExit()

	//start a thread to send heart beat to master
	go cluster.NodeLiveHeartBeat(cluster.GMasterPort, cluster.GNodeName)

	// start a rpc server for master and to access
	node := new(cluster.Node)
	rpc.Register(node)
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", "127.0.0.1:"+nodeport)

	if e != nil {
		log.Fatal("node listen error:", e)
	}

	log.Printf("%s is Starting......", cluster.GNodeName)
	http.Serve(l, nil)

}
