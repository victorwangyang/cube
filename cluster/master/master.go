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

	cluster.GMasterInfo.MasterName = os.Args[1]
	cluster.GMasterInfo.MasterPort = os.Args[2]
	cluster.GMasterInfo.NodeNumber = os.Args[3]

	cluster.GNodeInfo = make(map[string]cluster.NodeInfo)

	//start nodes here,so master can save all informations of nodes
	cluster.APIV1StartNodesDeamon(cluster.GMasterInfo.NodeNumber, cluster.GMasterInfo.MasterPort)

	//start a thread to watch the live state of master process
	go cluster.MasterExit()

	// start a rpc server for doserver and nodes to access
	master := new(cluster.Master)
	rpc.Register(master)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", "127.0.0.1:"+cluster.GMasterInfo.MasterPort)

	if e != nil {
		log.Fatal("master listen error:", e)
	}

	log.Printf("%s is Starting......", cluster.GMasterInfo.MasterName)

	http.Serve(l, nil)
}
