package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	// Flags
	var translateOrder int
	var libreTranslate bool
	var libreTranslateTarget string
	var camTranslateDictionary string
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
				printForRofi(
					args[0],
					translateOrder,
					libreTranslate,
					libreTranslateTarget,
					camTranslateDictionary,
				)
			}
		},
	}
	rootCmd.AddCommand(translateCmd)
	rootCmd.CompletionOptions.HiddenDefaultCmd = true
	translateCmd.Flags().IntVarP(&translateOrder, "order", "o", 1, "Order for translation")
	translateCmd.Flags().BoolVarP(&libreTranslate, "libre", "l", false, "Use libre translate or not")
	translateCmd.Flags().StringVarP(&libreTranslateTarget, "libre_target", "t", "en", "Target language for libretranslate")
	translateCmd.Flags().StringVarP(&camTranslateDictionary, "cam_dict", "d", "english", "Target language for cambridge dictionary")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
