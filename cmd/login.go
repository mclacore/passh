package cmd

import (
	"fmt"

	"github.com/mclacore/passh/pkg/password"
	"github.com/mclacore/passh/pkg/login"
	"github.com/spf13/cobra"
)

func NewCmdLogin() *cobra.Command {
	loginCmd := &cobra.Command{
		Use:   "login",
		Short: "Create, update, or retreive a login credential",
	}

	var login LoginItem

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

func runNewLogin(input *LoginItem) error {
	if input.password == "" {
		input.password = password.GeneratePassword(12, false, true, true, true)
	}

	loginItem := [][]string{
		{input.itemName},
		{input.username},
		{input.password},
		{input.url},
	}


	// loginFile creates the file
	fmt.Printf("login item: %v\n", loginItem)
	// loginItem is creating items
	// loginItems values are not being written to the file

	// insert sqlite shit here

	return nil
}
