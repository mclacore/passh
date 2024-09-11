package cmd

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/spf13/cobra"
	// "github.com/mclacore/passh/pkg/login"
)

type loginItem struct {
	itemName string
	username string
	password string
	url      string
}

func NewCmdLogin() *cobra.Command {
	loginCmd := &cobra.Command{
		Use:   "login",
		Short: "Create, update, or retreive a login credential",
	}

	var login loginItem

	loginNewCmd := &cobra.Command{
		Use:   "new",
		Short: "Create a new login credential",
		Run:   func(cmd *cobra.Command, args []string) { loginNewCmd(&login) },
	}

	genPass := runNewPass(passNewCmd, "--uppercase, --numbers, --special") // need to get the result of this function call to pass in a default.
	loginNewCmd.Flags().StringVarP(&login.itemName, "item-name", "i", "", "Name for the login item")
	loginNewCmd.Flags().StringVarP(&login.username, "username", "u", "", "Username for the login credential")
	loginNewCmd.Flags().StringVarP(&login.password, "password", "p", genPass, "Password for the login credential")
	loginNewCmd.Flags().StringVarP(&login.url, "url", "r", "", "URL for the login credential")

	loginCmd.AddCommand(loginNewCmd)
	return loginCmd
}

func loginNewCmd(input *loginItem) {
	loginItem := [][]string{
		{input.itemName},
		{input.username},
		{input.password},
		{input.url},
	}

	// add check for if file exists
	loginFile, err := os.Create("temp.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer loginFile.Close()

	writer := csv.NewWriter(loginFile)

	defer writer.Flush()

	writer.WriteAll(loginItem)
}
