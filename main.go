package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kr/pretty"
)

func print(w http.ResponseWriter, req *http.Request) {
	// Read body
	b, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var obj interface{}
	json.Unmarshal(b, &obj)

	log.Print("======================================== HEADER ================================================")
	pretty.Log(req.Header)
	log.Print("======================================== END HEADER ================================================")
	log.Print("=========================================== PAYLOAD =============================================")
	pretty.Log(obj)
	log.Print("=========================================== END PAYLOAD =============================================")

	fmt.Fprintf(w, "OK\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	http.HandleFunc("/", print)

	http.ListenAndServe(":80", nil)
}
