package cluster

import (
	"log"
	"net/http"
	"os"
	"time"
)

//GNodelive is the flag of liveness
var GNodelive = true

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

//NodeProcessRequestFromMaster is to Process request from Client
func NodeProcessRequestFromMaster(r *http.Request) {

	log.Println("NodeProcessRequestFromMaster .......")
}
