package net

import (
	"fmt"
	"github.com/jakenotjacob/neptr/internal/plugin"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	pattern := "/status"
	handler := func(response http.ResponseWriter, request *http.Request) {
		io.WriteString(response, "Reached status handler!\n")
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			log.Printf("Unable to read response body")
		}
		defer request.Body.Close()

		log.Printf("Body: %v\n", string(body))
		//log.Println(request.Header["name"])
	}
	http.HandleFunc(pattern, handler)

	action := plugin.New("/foo", func(w http.ResponseWriter, r *http.Request) { log.Printf("Running callback for: foo\n") })
	action.Register()

	actionStatusHandler := func(response http.ResponseWriter, request *http.Request) {
		io.WriteString(response, fmt.Sprintf("%v\n", action.State))
	}
	http.HandleFunc("/foo/status", actionStatusHandler)

	actionRunHandler := func(response http.ResponseWriter, request *http.Request) {
		log.Println("Running %v\n", action.Name)
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			log.Printf("Unable to read response body")
		}
		defer request.Body.Close()

		//var a Action
		//a, err := json.Unmarshal(&a, body)
		//if err != nil {
		//	// Wasnt request body we could understand
		//}
		// go action.Exec(a.Fn)

		if len(body) > 0 {
			go action.Exec(string(body))
		}
		//action.Run()
	}
	http.HandleFunc("/foo/run", actionRunHandler)

	log.Printf("Hello log")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
