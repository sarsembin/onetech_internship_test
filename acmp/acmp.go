package acmp

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func Difficulty(url string) float64 {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	strBody := string(body)

	diffIndexStart := strings.Index(strBody, "<i>")
	diffIndexEnd := strings.Index(strBody, "</i>")
	if diffIndexStart == -1 {
		return -1
	}

	strDiff := strBody[diffIndexStart:diffIndexEnd]
	diffIndex := strings.Index(strDiff, "%")

	diffString := strings.TrimSpace(strDiff[diffIndex-2:diffIndex])
	difficulty, err := strconv.ParseFloat(diffString, 64)
	if err != nil {
		log.Fatal(err)
	}

	return difficulty
}



