package cmd

import (
	"fmt"

	"github.com/mclacore/passh/pkg/login"
	"github.com/mclacore/passh/pkg/password"
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
		RunE:  runNewLogin,
	}
	loginNewCmd.MarkFlagRequired("item-name")
	loginNewCmd.Flags().StringP("username", "u", "", "Username for the login credential")
	loginNewCmd.Flags().StringP("password", "p", "", "Password for the login credential")
	loginNewCmd.Flags().StringP("url", "r", "", "URL for the login credential")

	loginGetCmd := &cobra.Command{
		Use:   "get",
		Short: "Get login item properties",
		RunE:  runGetLogin,
	}
	loginGetCmd.MarkFlagsOneRequired()

	loginUpdateCmd := &cobra.Command{
		Use:   "update",
		Short: "Update login item property",
		RunE:  runUpdateLogin,
	}
	loginUpdateCmd.MarkFlagRequired("item-name")
	loginUpdateCmd.Flags().StringP("username", "u", "", "Username to update")
	loginUpdateCmd.Flags().StringP("password", "p", "", "Password to update")
	loginUpdateCmd.Flags().StringP("url", "r", "", "URL to update")

	loginCmd.AddCommand(loginNewCmd)
	loginCmd.AddCommand(loginGetCmd)
	loginCmd.AddCommand(loginUpdateCmd)
	loginCmd.PersistentFlags().StringP("item-name", "i", "", "Name for the login item")
	return loginCmd
}

func runNewLogin(cmd *cobra.Command, args []string) error {
	itemName, itemErr := cmd.Flags().GetString("item-name")
	if itemErr != nil {
		return fmt.Errorf("Error setting item-name: %v", itemErr)
	}

	username, userErr := cmd.Flags().GetString("username")
	if userErr != nil {
		return fmt.Errorf("Error setting username: %v", userErr)
	}

	loginPass, passErr := cmd.Flags().GetString("password")
	if passErr != nil {
		return fmt.Errorf("Error setting password: %v", passErr)
	}
	if loginPass == "" {
		loginPass = password.GeneratePassword(12, false, true, true, true)
	}

	url, urlErr := cmd.Flags().GetString("url")
	if urlErr != nil {
		return fmt.Errorf("Error setting URL: %v", urlErr)
	}

	loginItem := login.LoginItem{
		LoginItem: itemName,
		Username:  username,
		Password:  loginPass,
		URL:       url,
	}

	db, dbErr := login.ConnectToDB()
	if dbErr != nil {
		return fmt.Errorf("Error connecting to database: %V", dbErr)
	}

	createErr := login.CreateLoginItem(db, loginItem)
	if createErr != nil {
		return fmt.Errorf("Error creating new login item: %v", createErr)
	}

	return nil
}

func runGetLogin(cmd *cobra.Command, args []string) error {
	itemName, itemErr := cmd.Flags().GetString("item-name")
	if itemErr != nil {
		return fmt.Errorf("Error setting item-name: %v", itemErr)
	}

	db, dbErr := login.ConnectToDB()
	if dbErr != nil {
		return fmt.Errorf("Error connecting to database: %v", dbErr)
	}

	getItem, getErr := login.GetLoginItem(db, itemName)
	if getErr != nil {
		return fmt.Errorf("Error fetching login item: %v", getErr)
	}

	cmd.Printf("Login item: %v\n", getItem.LoginItem)
	cmd.Printf("Username: %v\n", getItem.Username)
	cmd.Printf("Password: %v\n", getItem.Password)
	cmd.Printf("URL: %v\n", getItem.URL)

	return nil
}

func runUpdateLogin(cmd *cobra.Command, args []string) error {
	itemName, itemErr := cmd.Flags().GetString("item-name")
	if itemErr != nil {
		return fmt.Errorf("Error setting item-name: %v", itemErr)
	}

	db, dbErr := login.ConnectToDB()
	if dbErr != nil {
		return fmt.Errorf("Error connecting to database: %v", dbErr)
	}

	newLoginItem, newLoginErr := login.GetLoginItem(db, itemName)
	if newLoginErr != nil {
		return fmt.Errorf("Error fetching login item to update: %v", newLoginErr)
	}

	if username, userErr := cmd.Flags().GetString("username"); userErr != nil {
		newLoginItem.Username = username
	}

	if password, passErr := cmd.Flags().GetString("password"); passErr != nil {
		newLoginItem.Password = password
	}

	if url, urlErr := cmd.Flags().GetString("url"); urlErr != nil {
		newLoginItem.URL = url
	}

	login.UpdateLoginItem(db, newLoginItem)
	runGetLogin(cmd, args)

	return nil
}
