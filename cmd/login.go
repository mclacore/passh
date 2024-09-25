package cmd

import (
	"github.com/mclacore/passh/pkg/password"
	"github.com/mclacore/passh/pkg/login"
	"github.com/spf13/cobra"
)

func NewCmdLogin() *cobra.Command {
	loginCmd := &cobra.Command{
		Use:   "login",
		Short: "Create, update, or retreive a login credential",
	}

	loginNewCmd := &cobra.Command{
		Use:   "new",
		Short: "Create a new login credential",
		RunE: runNewLogin,
	}

	loginNewCmd.Flags().StringP("item-name", "i", "", "Name for the login item")
	_ = loginNewCmd.MarkFlagRequired("item-name")
	loginNewCmd.Flags().StringP("username", "u", "", "Username for the login credential")
	loginNewCmd.Flags().StringP("password", "p", "", "Password for the login credential")
	loginNewCmd.Flags().StringP("url", "r", "", "URL for the login credential")

	loginCmd.AddCommand(loginNewCmd)
	return loginCmd
}

func runNewLogin(cmd *cobra.Command, args []string) error {
	itemName, itemErr := cmd.Flags().GetString("item-name")
	if itemErr != nil {
		return itemErr
	}

	username, userErr := cmd.Flags().GetString("username")
	if userErr != nil {
		return userErr
	}

	loginPass, passErr := cmd.Flags().GetString("password")
	if passErr != nil {
		return passErr
	}
	if loginPass == "" {
		loginPass = password.GeneratePassword(12, false, true, true, true)
	}

	url, urlErr := cmd.Flags().GetString("url")
	if urlErr != nil {
		return urlErr
	}

	loginItem := login.LoginItem{
		LoginItem: itemName,
		Username: username,
		Password: loginPass,
		URL: url,
	}

	db, dbErr := login.ConnectToDB()
	if dbErr != nil {
		return dbErr
	}

	createErr := login.CreateLoginItem(db, loginItem)
	if createErr != nil {
		return createErr
	}

	return nil
}
