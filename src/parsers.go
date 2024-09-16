package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type ParsedResponse struct {
	Translate   string
	Description string
}

func parseCam(word string, dict string) []string {
	url := fmt.Sprintf("https://dictionary.cambridge.org/dictionary/%s/%s", dict, word)
	headers := make(map[string]string)
	headers["Content-Type"] = "json"
	resp := getRequest(url, headers)
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var result []string
	for _, value := range doc.Find(".def.ddef_d.db").EachIter() {
		str := strings.ToLower(strings.ReplaceAll(value.Text(), "\n", ""))
		str = strings.ReplaceAll(str, "\t", "")
		str = strings.ReplaceAll(str, ":", "")
		splitted := strings.Fields(str) // That is ugly man
		var output string
		for _, substring := range splitted {
			output += substring + " "
		}
		result = append(result, upperFirstLetter(output))
	}
	return result
}

type LibreTranslateRespone struct {
	Alternatives   []string `json:"alternatives"`
	TranslatedText string   `json:"translatedText"`
}

func parseLibreTranslate(word string, target string) []string {
	url := "http://localhost:5000/translate"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	body := []byte(fmt.Sprintf(`{ "source": "%s", "target": "%s", "q": "%s", "alternatives": 3 }`, target, "ru", word))
	var result LibreTranslateRespone
	err, resp := postRequest(url, headers, body)
	if err != nil {
		return []string{}
	}
	defer resp.Body.Close()
	body, nerr := io.ReadAll(resp.Body)
	if nerr != nil {
		log.Panic(nerr)
	}
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}
	var output []string
	output = append(output, result.TranslatedText)
	for _, alternative := range result.Alternatives {
		output = append(output, alternative)
	}
	return output
}
