package main

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
	"io/ioutil"
)
type item struct{
	Name string `json:"num"`
}
type rest struct {
	Id int `json:"id"`
	Model string `json:"model"`
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "hello")
		if err!=nil {
			log.Panic(err)
		}
}
func sendResp(w http.ResponseWriter, r *http.Request) {
	// decoder := json.NewDecoder(r.Body)
	body, err := ioutil.ReadAll(r.Body)
	if err!=nil {
		fmt.Println("the err in ioutil")
		panic(err)
	}
	var req item
	err = json.Unmarshal(body, &req)
	if err!=nil {
		fmt.Println("err in decoding json")
	}
	fmt.Println(req)
	var re rest
	re.Id = 105
	re.Model = "Lenevo"
	json, err := json.Marshal(re)
	fmt.Println(w, string(json))
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)

}
func main() {
	fs := http.FileServer(http.Dir("../forms/dist"))

	http.Handle("/", fs)
	http.HandleFunc("/items", sendResp)
	fmt.Println("Server listening at port 8000")
	log.Panic(http.ListenAndServe(":8000", nil))
}