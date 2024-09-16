package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
)

type ParsedResponse struct {
	Translate   string
	Description string
}

func parseCam(word string) []string {
	url := "https://dictionary.cambridge.org/dictionary/english/" + word
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

func parseLibreTranslate(word string) []string {
	url := "http://localhost:5000/translate"
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	body := []byte(fmt.Sprintf(`{ "source": "%s", "target": "%s", "q": "%s", "alternatives": 3 }`, "en", "ru", word))
	var result LibreTranslateRespone
	resp := postRequest(url, headers, body)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
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

func printForRofi(word string, translateOrdered bool) {
	libreResult := parseLibreTranslate(word)
	cambrResult := parseCam(word)
	splitResult := []string{strings.Repeat("-", 100)}
	mergedResult := append(cambrResult, splitResult...)
	mergedResult = append(mergedResult, libreResult...)
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

func main() {
	var rootCmd = &cobra.Command{
		Use:   "ttr",
		Short: "Rofi translate cli app",
		Long:  `Entry point for rofi-translate app`,

		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("No action provided. Please run cli with --help flag.")
		},
	}
	var translateCmd = &cobra.Command{
		Use:   "translate",
		Short: "Translate word or sentence",
		Long:  `Really long string`,

		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				printForRofi(args[0], true)
			}
		},
	}
	rootCmd.AddCommand(translateCmd)
	rootCmd.CompletionOptions.HiddenDefaultCmd = true
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
