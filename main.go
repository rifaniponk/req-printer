package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/kr/pretty"
	logr "github.com/sirupsen/logrus"
)

var (
	PRETTY_PAYLOAD = true
	PRETTY_HEADER  = true
)

func init() {
	PRETTY_PAYLOAD = os.Getenv("PRETTY_PAYLOAD") == "true"
	PRETTY_HEADER = os.Getenv("PRETTY_HEADER") == "true"
}

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
	if PRETTY_PAYLOAD {
		pretty.Log(req.Header)
	} else {
		logr.WithFields(logr.Fields{"header": req.Header}).Info()
	}
	log.Print("======================================== END HEADER ================================================")

	log.Print("=========================================== PAYLOAD =============================================")
	if PRETTY_PAYLOAD {
		pretty.Log(obj)
	} else {
		log.Printf(string(b))
	}
	log.Print("=========================================== END PAYLOAD =============================================")

	fmt.Fprintf(w, "OK\n")
}

func main() {

	http.HandleFunc("/", print)

	http.ListenAndServe(":80", nil)
}
