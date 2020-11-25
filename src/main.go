package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func GetFITJByPageNumber(page int) FemicidesInTurkeyJson {
retry:
	url := "https://femicidesinturkey.com/api/victim/?page=" + strconv.Itoa(page)
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}

	if resp.StatusCode == 200 {
		var result FemicidesInTurkeyJson
		responseBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
		}
		json.Unmarshal(responseBody, &result)
		return result
	}

	goto retry
}

func main() {
	maxPage := GetFITJByPageNumber(1).Information.Pages

	for i := 1; i < maxPage; i++ {
		var FITJ FemicidesInTurkeyJson = GetFITJByPageNumber(i)
		for _, v := range FITJ.Data {
			println(v.CreateSentence())
		}
	}
}
