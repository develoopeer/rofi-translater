package main

import (
	"fmt"
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

func parseCam(word string) {
	url := "https://dictionary.cambridge.org/dictionary/english/" + word
	headers := make(map[string]string)
	headers["Content-Type"] = "json"
	resp := makeRequest(url, headers)
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	prompt_string := fmt.Sprintf("\000prompt\x1f%s\n", word)
	var count int
	fmt.Print(prompt_string)
	for i, value := range doc.Find(".def.ddef_d.db").EachIter() {
		str := strings.ToLower(strings.ReplaceAll(value.Text(), "\n", ""))
		str = strings.ReplaceAll(str, "\t", "")
		str = strings.ReplaceAll(str, ":", "")
		splitted := strings.Fields(str) // That is ugly man
		var output string
		for _, substring := range splitted {
			output += substring + " "
		}
		fmt.Println(upperFirstLetter(output))
		count = i + 1 // i starts with 0
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
				parseCam(args[0])
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
