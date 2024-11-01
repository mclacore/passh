package cmd

import (
	"fmt"

	"github.com/mclacore/passh/pkg/collection"
	"github.com/mclacore/passh/pkg/database"
	"github.com/mclacore/passh/pkg/login"
	"github.com/mclacore/passh/pkg/password"
	"github.com/mclacore/passh/pkg/prompt"
	"github.com/spf13/cobra"
)

func NewCmdLogin() *cobra.Command {
	var login login.LoginItem

	loginCmd := &cobra.Command{
		Use:   "login",
		Short: "Create, update, or retreive a login credential",
	}

	loginNewCmd := &cobra.Command{
		Use:   "new",
		Short: "Create a new login credential",
		RunE:  runNewLogin,
	}
	loginNewCmd.MarkFlagRequired("collection-name")
	loginNewCmd.MarkFlagRequired("item-name")
	loginNewCmd.Flags().StringP("username", "u", "", "Username for the login credential")
	loginNewCmd.Flags().StringP("password", "p", "", "Password for the login credential")
	loginNewCmd.Flags().StringP("url", "r", "", "URL for the login credential")
	loginNewCmd.Flags().Bool("no-password", false, "Do not set a password for login")

	loginGetCmd := &cobra.Command{
		Use:   "get",
		Short: "Get login item properties",
		RunE:  runGetLogin,
	}
	loginGetCmd.MarkFlagRequired("collection-name")
	loginGetCmd.MarkFlagRequired("item-name")
	loginGetCmd.Flags().BoolP("show-password", "p", false, "Show password")

	loginUpdateCmd := &cobra.Command{
		Use:   "update",
		Short: "Update login item property",
		RunE:  runUpdateLogin,
	}
	loginUpdateCmd.MarkFlagRequired("collection-name")
	loginUpdateCmd.MarkFlagRequired("item-name")
	loginUpdateCmd.Flags().StringP("username", "u", login.Username, "Username to update")
	loginUpdateCmd.Flags().StringP("password", "p", login.Password, "Password to update")
	loginUpdateCmd.Flags().StringP("url", "r", login.URL, "URL to update")

	loginListCmd := &cobra.Command{
		Use:   "list",
		Short: "List all login items",
		RunE:  runListLogins,
	}
	loginListCmd.MarkFlagRequired("collection-name")

	loginDeleteCmd := &cobra.Command{
		Use:     "delete",
		Aliases: []string{"del, remove, rm"},
		Short:   "Delete a login item",
		RunE:    runDeleteLogin,
	}
	loginDeleteCmd.MarkFlagRequired("item-name")
	loginDeleteCmd.MarkFlagRequired("collection-name")

	loginCmd.AddCommand(loginNewCmd)
	loginCmd.AddCommand(loginGetCmd)
	loginCmd.AddCommand(loginUpdateCmd)
	loginCmd.AddCommand(loginListCmd)
	loginCmd.AddCommand(loginDeleteCmd)
	loginCmd.PersistentFlags().StringP("item-name", "i", "", "Name for the login item")
	loginCmd.PersistentFlags().StringP("collection-name", "c", "default", "Name for the login collection")
	return loginCmd
}

func runNewLogin(cmd *cobra.Command, args []string) error {
	itemName, itemErr := cmd.Flags().GetString("item-name")
	if itemErr != nil {
		return fmt.Errorf("Error setting item-name: %w", itemErr)
	}

	username, userErr := cmd.Flags().GetString("username")
	if userErr != nil {
		return fmt.Errorf("Error setting username: %w", userErr)
	}

	loginPass, passErr := cmd.Flags().GetString("password")
	if passErr != nil {
		return fmt.Errorf("Error setting password: %w", passErr)
	}

	noPass, noPassErr := cmd.Flags().GetBool("no-password")
	if noPassErr != nil {
		return fmt.Errorf("Error skipping password: %w", noPassErr)
	}

	if loginPass == "" && !noPass {
		loginPass = password.GeneratePassword(12, false, true, true, true)
	}

	url, urlErr := cmd.Flags().GetString("url")
	if urlErr != nil {
		return fmt.Errorf("Error setting URL: %w", urlErr)
	}

	col, colErr := cmd.Flags().GetString("collection-name")
	if col == "" {
		col = "default"
	}
	if colErr != nil {
		return fmt.Errorf("Error setting collection: %w", colErr)
	}

	db, dbErr := database.ConnectToDB()
	if dbErr != nil {
		return fmt.Errorf("Error connecting to database: %w", dbErr)
	}

	colId, colIdErr := collection.GetCollection(db, col)
	if colId == nil {
		defCol := collection.Collection{
			Name: "default",
		}
		collection.CreateCollection(db, defCol)
		colId, colIdErr = collection.GetCollection(db, "default")
	}
	if colIdErr != nil {
		return fmt.Errorf("Error getting collection id: %w", colIdErr)
	}

	loginItem := login.LoginItem{
		ItemName:     itemName,
		Username:     username,
		Password:     loginPass,
		URL:          url,
		CollectionID: int(colId.ID),
	}

	createErr := login.CreateLoginItem(db, loginItem)
	if createErr != nil {
		return fmt.Errorf("Error creating new login item: %w", createErr)
	}

	return nil
}

func runGetLogin(cmd *cobra.Command, args []string) error {
	itemName, itemErr := cmd.Flags().GetString("item-name")
	if itemErr != nil {
		return fmt.Errorf("Error setting item-name: %w", itemErr)
	}

	db, dbErr := database.ConnectToDB()
	if dbErr != nil {
		return fmt.Errorf("Error connecting to database: %w", dbErr)
	}

	// need to do collection checking here
	getItem, getErr := login.GetLoginItem(db, itemName)
	if getErr != nil {
		return fmt.Errorf("Error fetching login item: %w", getErr)
	}

	showPass, passErr := cmd.Flags().GetBool("show-password")
	if passErr != nil {
		return fmt.Errorf("Error showing password: %w", passErr)
	}

	// need to add a for loop here for all listed item-names that match itemName

	cmd.Printf("Item Name: %v\n", getItem.ItemName)
	cmd.Printf("Username: %v\n", getItem.Username)

	if showPass {
		cmd.Printf("Password: %v\n", getItem.Password)
	} else {
		cmd.Println("Password: <hidden>")
	}

	cmd.Printf("URL: %v\n", getItem.URL)

	return nil
}

func runUpdateLogin(cmd *cobra.Command, args []string) error {
	itemName, itemErr := cmd.Flags().GetString("item-name")
	if itemErr != nil {
		return fmt.Errorf("Error setting item-name: %w", itemErr)
	}

	db, dbErr := database.ConnectToDB()
	if dbErr != nil {
		return fmt.Errorf("Error connecting to database: %w", dbErr)
	}

	newLoginItem, newLoginErr := login.GetLoginItem(db, itemName)
	if newLoginErr != nil {
		return fmt.Errorf("Error fetching login item to update: %w", newLoginErr)
	}

	username, userErr := cmd.Flags().GetString("username")
	if userErr != nil {
		return fmt.Errorf("Error updating username: %w", userErr)
	}

	password, passErr := cmd.Flags().GetString("password")
	if passErr != nil {
		return fmt.Errorf("Error updating password: %w", passErr)
	}

	url, urlErr := cmd.Flags().GetString("url")
	if urlErr != nil {
		return fmt.Errorf("Error updating URL: %w", urlErr)
	}

	if len(username) > 0 {
		newLoginItem.Username = username
	}
	if len(password) > 0 {
		newLoginItem.Password = password
	}
	if len(url) > 0 {
		newLoginItem.URL = url
	}

	login.UpdateLoginItem(db, newLoginItem)
	runGetLogin(cmd, args)

	return nil
}

func runListLogins(cmd *cobra.Command, args []string) error {
	db, dbErr := database.ConnectToDB()
	if dbErr != nil {
		return fmt.Errorf("Error connecting to database: %w", dbErr)
	}

	// need to be able to filter specifically from collections here
	items, itemErr := login.ListLoginItems(db)
	if itemErr != nil {
		return fmt.Errorf("Error fetching login items: %w", itemErr)
	}

	for _, item := range *items {
		cmd.Println(item.ItemName)
	}
	return nil
}

func runDeleteLogin(cmd *cobra.Command, args []string) error {
	db, dbErr := database.ConnectToDB()
	if dbErr != nil {
		return fmt.Errorf("Error connecting to database: %w", dbErr)
	}

	itemToDel, itemDelErr := cmd.Flags().GetString("item-name")
	if itemDelErr != nil {
		return fmt.Errorf("Error finding item to delete: %w", itemDelErr)
	}

	confirm, confirmErr := prompt.ConfirmItemDelete()
	if confirmErr != nil {
		cmd.Printf("Operation cancelled.\n")
		return nil
	}

	if confirm == "y" || confirm == "Y" {
		delErr := login.DeleteLoginItem(db, itemToDel)
		if delErr != nil {
			return fmt.Errorf("Error deleting item: %w", delErr)
		}
		cmd.Printf("%v has been deleted.\n", itemToDel)
	}

	return nil
}
