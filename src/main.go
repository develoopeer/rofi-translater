package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func main() {
	// Flags
	var splitString int
	var translateOrder int
	var libreTranslate bool
	var libreTranslateTarget string
	var camTranslateDictionary string

	var exportWord string
	var exportMeaning string
	var exportTranslate string

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
			if os.Getenv("ROFI_RETV") == "1" { // Entry entered
				appendToCsv(
					"/home/laptop/.config/ttr/output.csv",
					[]string{os.Getenv("ROFI_DATA"), args[0]},
				)
				exec.Command("notify-send", "New word have been added", os.Getenv("ROFI_DATA")).Run()
				printForRofi(
					os.Getenv("ROFI_DATA"),
					translateOrder,
					libreTranslate,
					libreTranslateTarget,
					camTranslateDictionary,
				)
			} else {
				if len(args) > 0 {
					printForRofi(
						args[0],
						translateOrder,
						libreTranslate,
						libreTranslateTarget,
						camTranslateDictionary,
					)
				}
			}
		},
	}

	var exportCmd = &cobra.Command{
		Use:   "export",
		Short: "Add a word to flashcards",

		Run: func(cmd *cobra.Command, args []string) {
			appendToCsv(
				args[0],
				[]string{
					exportWord,
					exportMeaning,
					exportTranslate,
				},
			)

		},
	}
	rootCmd.AddCommand(translateCmd)
	rootCmd.AddCommand(exportCmd)
	rootCmd.CompletionOptions.HiddenDefaultCmd = true

	exportCmd.Flags().StringVarP(&exportWord, "word", "w", "", "Word column")
	exportCmd.Flags().StringVarP(&exportMeaning, "meaning", "m", "", "Meaning column")
	exportCmd.Flags().StringVarP(&exportTranslate, "translation", "t", "", "Translation column")

	translateCmd.Flags().IntVarP(&splitString, "split", "s", 1, "Split string")
	translateCmd.Flags().IntVarP(&translateOrder, "order", "o", 1, "Order for translation")
	translateCmd.Flags().BoolVarP(&libreTranslate, "libre", "l", false, "Use libre translate or not")
	translateCmd.Flags().StringVarP(&libreTranslateTarget, "libre_target", "t", "en", "Target language for libretranslate")
	translateCmd.Flags().StringVarP(&camTranslateDictionary, "cam_dict", "d", "english", "Target language for cambridge dictionary")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
