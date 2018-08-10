package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	testGet("f1")
	testPost()
}

func testGet(p string) {
	url2 := "http://127.0.0.1:7788/amount/" + p

	req, _ := http.NewRequest("GET", url2, nil)

	client := &http.Client{}
	resp, _ := client.Do(req)

	body, _ := ioutil.ReadAll(resp.Body)

	var result string

	json.Unmarshal(body, &result)

	fmt.Printf("%s\n", string(body[:]))

	defer resp.Body.Close()
}

func testPost() {
	url2 := "http://127.0.0.1:7788/basic"

	//v := url.Values{}
	//v.Set("name", "xxxx")
	//v.Set("email", "xxxx")
	//data := v.Encode()
	payload := strings.NewReader("name=cute&email=test@a.a")
	req, _ := http.NewRequest("POST", url2, payload)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;")

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var result string
	json.Unmarshal(body, &result)
	fmt.Printf("%s\n", string(body[:]))

}
