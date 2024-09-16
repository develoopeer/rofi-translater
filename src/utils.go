package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func getRequest(url string, headers map[string]string) *http.Response {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Panic(err)
	}
	req.Header.Add("User-Agent", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36`)
	for key, value := range headers {
		req.Header.Add(key, value)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}

	return resp
}

func postRequest(url string, headers map[string]string, body []byte) (error, *http.Response) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		log.Panic(err)
	}
	req.Header.Add("User-Agent", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36`)
	for key, value := range headers {
		req.Header.Add(key, value)
	}
	resp, err := client.Do(req)
	return err, resp
}

func upperFirstLetter(s string) string {
	return strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
}

func printForRofi(
	word string,
	translateOrder int,
	libreTranslate bool,
	libreTranslateTarget string,
	camTranslateDict string,
) {
	var libreResult []string
	var mergedResult []string
	if libreTranslate == false {
		libreResult = []string{}
	} else {
		libreResult = parseLibreTranslate(word, libreTranslateTarget)
	}
	cambrResult := parseCam(word, camTranslateDict)
	splitResult := []string{strings.Repeat("-", 100)}
	if translateOrder == 1 {
		mergedResult = append(cambrResult, splitResult...)
		mergedResult = append(mergedResult, libreResult...)
	}
	if translateOrder == 2 {
		mergedResult = append(libreResult, splitResult...)
		mergedResult = append(mergedResult, cambrResult...)
	}
	prompt_string := fmt.Sprintf("\000prompt\x1f%s\n", word)
	fmt.Print(prompt_string)
	var count int
	for i, line := range mergedResult {
		fmt.Println(line)
		count = i + 1
	}
	if count == 0 {
		fmt.Println("No output")
	}

}
