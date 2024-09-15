package cmd

import (
	"encoding/csv"
	"errors"
	"os"

	"github.com/mclacore/passh/pkg/password"
	"github.com/spf13/cobra"
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
		Run:   func(cmd *cobra.Command, args []string) { runNewLogin(&login) },
	}

	loginNewCmd.Flags().StringVarP(&login.itemName, "item-name", "i", "", "Name for the login item")
	_ = loginNewCmd.MarkFlagRequired("item-name")
	loginNewCmd.Flags().StringVarP(&login.username, "username", "u", "", "Username for the login credential")
	loginNewCmd.Flags().StringVarP(&login.password, "password", "p", "", "Password for the login credential")
	loginNewCmd.Flags().StringVarP(&login.url, "url", "r", "", "URL for the login credential")

	loginCmd.AddCommand(loginNewCmd)
	return loginCmd
}

func runNewLogin(input *loginItem) error {
	if input.password == "" {
		input.password = password.GeneratePassword(12, false, true, true, true)
	}

	loginItem := [][]string{
		{input.itemName},
		{input.username},
		{input.password},
		{input.url},
	}

	if _, err := os.Stat("temp.csv"); errors.Is(err, os.ErrNotExist) {
		loginFile, createErr := os.Create("temp.csv")
		if createErr != nil {
			return createErr
		}
	}

	// add check for if file exists
	defer loginFile.Close()

	writer := csv.NewWriter(loginFile)

	defer writer.Flush()

	return writer.WriteAll(loginItem)
}
