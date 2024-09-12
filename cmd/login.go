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

	loginNewCmd.Flags().StringVarP(&login.itemName, "item-name", "i", "", "Name for the login item")
	loginNewCmd.Flags().StringVarP(&login.username, "username", "u", "", "Username for the login credential")
	loginNewCmd.Flags().StringVarP(&login.password, "password", "p", "", "Password for the login credential")
	loginNewCmd.Flags().StringVarP(&login.url, "url", "r", "", "URL for the login credential")

	loginCmd.AddCommand(loginNewCmd)
	return loginCmd
}

func loginNewCmd(cmd *cobra.Command, input *loginItem) error {
	var password string

	passwordArg, argErr := cmd.Flags().GetString("password")
	if argErr != nil {
		return argErr
	}

	if passwordArg == "" {
		args := []string{"--uppercase", "--numbers", "--special"}

		password = string(runNewPass(cmd, args))
	} else {
		password = input.password
	}

	loginItem := [][]string{
		{input.itemName},
		{input.username},
		{password},
		{input.url},
	}

	// add check for if file exists
	loginFile, createErr := os.Create("temp.csv")
	if createErr != nil {
		return createErr
	}
	defer loginFile.Close()

	writer := csv.NewWriter(loginFile)

	defer writer.Flush()

	return writer.WriteAll(loginItem)
}
