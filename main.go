package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	url := "https://api.exmo.com/v1/ticker/"

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	type Crypto map[string]map[string]interface{}

	var f Crypto

	jsonErr := json.Unmarshal(body, &f)

	if jsonErr != nil {
		log.Fatal(err)
	}

	for key, value := range f {
		for tags := range value {
			str := fmt.Sprintf("Trader,order=%v %v=%v", key, tags, value[tags])
			fmt.Println(str)
		}
	}
}
