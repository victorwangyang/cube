package main

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// Get is getting data from server by Restful API
func Get(url string, data []byte) (resp []byte) {

	ret, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	defer ret.Body.Close()

	body, err := ioutil.ReadAll(ret.Body)
	if err != nil {
		panic(err)
	}

	return body

}

// Post is sending data to server by Restful API
func Post(url string, data []byte) {

	resp, err := http.Post(url, "app", strings.NewReader(string(data)))

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

}

//Delete is deleting resource of cluster
func Delete(url string) {

	req, _ := http.NewRequest("DELETE", url, nil)
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

}
