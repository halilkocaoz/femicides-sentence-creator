package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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
	sentences := ""
	for i := 1; i < maxPage; i++ {
		var FITJ FemicidesInTurkeyJson = GetFITJByPageNumber(i)
		for _, v := range FITJ.Data {
			sentence, err := v.CreateSentence()
			if err == nil {
				sentences += sentence + "\n"
			}
		}
	}

	Output(sentences)
}

func Output(data string) error {
	file, err := os.Create("output.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
}
