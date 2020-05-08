package main

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

func httpRequest(url string) ([]byte, error) {
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		panic(err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
		return nil, err
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println("response" + string(body))
	return body, nil
}

func StatusUpdate() (Data,error) {
	url := "http://127.0.0.1:8080/UpdateData"
	dataAsByte, err := httpRequest(url)  //byte array
	if err != nil {
		log.Fatal(err.Error())
		return Data{},err
	}
	var getData Data       //สร้างตัวแปร result ด้วย struct Result
	err = json.Unmarshal(dataAsByte, &getData)
	if err != nil {
		log.Fatal(err.Error())
		return Data{},err
	}
	return getData, nil
}

type Data struct {
	Message string `json:"message"`
}

func main() {
	c := time.Tick(1 * time.Minute)
	status, err := StatusUpdate()
	if err != nil {
		log.Fatal(err.Error())
		panic(err.Error())
	}
	for next := range c {
		fmt.Printf("%v \n%s\n", next, status.Message)
	}
}

