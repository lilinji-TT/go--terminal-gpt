package cmd

import (
	"GTG/config"
	"GTG/model"
	"GTG/pkg/gpt"
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var chatCmd = &cobra.Command{
	Use: "chat",
}

var GlobMessages []model.Message

func Chat(cmd *cobra.Command, args []string) {

	if missConfig() {
		fmt.Println("Please set your config, url and api key. GTG config -u <your url> -k <your api key>")
		return
	}

	fmt.Println("Welcome To Use GoTerminalGPT")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println()
		fmt.Print("User ~ % ")
		scanner.Scan()
		fmt.Println()
		userInput := scanner.Text()
		if userInput == "exit" {
			break
		}

		if userInput == "model" {
			config.SetModelName()
			continue
		}

		fmt.Printf("%s ~ %% ", config.Model)
		gpt.GenerateStreamWithGPT(userInput, &GlobMessages, config.Model)
		fmt.Println()
	}
}

func missConfig() bool {
	url, key, err := config.ReadConfig()
	return err != nil || len(url) == 0 || len(key) == 0
}

func init() {
	rootCmd.AddCommand(chatCmd)
	chatCmd.Run = Chat
}
