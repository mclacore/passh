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
	loginCmd.PersistentFlags().StringP("collection-name", "c", "", "Name for the login collection")
	loginCmd.MarkPersistentFlagRequired("collection-name")

	loginNewCmd := &cobra.Command{
		Use:   "new",
		Short: "Create a new login credential",
		RunE:  runNewLogin,
	}
	loginNewCmd.Flags().StringP("item-name", "i", "", "Name for the login item")
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
	loginGetCmd.Flags().StringP("item-name", "i", "", "Name for the login item")
	loginGetCmd.MarkFlagRequired("item-name")
	loginGetCmd.Flags().BoolP("show-password", "p", false, "Show password")

	loginUpdateCmd := &cobra.Command{
		Use:   "update",
		Short: "Update login item property",
		RunE:  runUpdateLogin,
	}
	loginUpdateCmd.Flags().StringP("item-name", "i", "", "Name for the login item")
	loginUpdateCmd.MarkFlagRequired("item-name")
	loginUpdateCmd.Flags().StringP("username", "u", login.Username, "Username to update")
	loginUpdateCmd.Flags().StringP("password", "p", login.Password, "Password to update")
	loginUpdateCmd.Flags().StringP("url", "r", login.URL, "URL to update")
	loginUpdateCmd.Flags().StringP("move-collection", "m", login.Collection.Name, "Collection to move to")

	loginListCmd := &cobra.Command{
		Use:   "list",
		Short: "List all login items",
		RunE:  runListLogins,
	}

	loginDeleteCmd := &cobra.Command{
		Use:     "delete",
		Aliases: []string{"del, remove, rm"},
		Short:   "Delete a login item",
		RunE:    runDeleteLogin,
	}
	loginDeleteCmd.Flags().StringP("item-name", "i", "", "Name for the login item")
	loginDeleteCmd.MarkFlagRequired("item-name")

	loginCmd.AddCommand(loginNewCmd)
	loginCmd.AddCommand(loginGetCmd)
	loginCmd.AddCommand(loginUpdateCmd)
	loginCmd.AddCommand(loginListCmd)
	loginCmd.AddCommand(loginDeleteCmd)
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
	// need to somehow create a collections table here if a collection is specified but does not exist as a table
	if colErr != nil {
		return fmt.Errorf("Error setting collection: %w", colErr)
	}

	db, dbErr := database.ConnectToDB()
	if dbErr != nil {
		return fmt.Errorf("Error connecting to database: %w", dbErr)
	}

	colId, colIdErr := collection.GetCollectionByName(db, col)
	if colId == nil {
		defCol := collection.Collection{
			Name: "default",
		}
		collection.CreateCollection(db, defCol)
		colId, colIdErr = collection.GetCollectionByName(db, "default")
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

	// need to check for unique constraints and error out about duplicate entry
	// currently passing through SQL error
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

	colName, colNameErr := cmd.Flags().GetString("collection-name")
	// remove this part, and instead use promptui to check if a collection exists. if not, prompt user to create one on the spot.
	// if only 1 collection exists, set it as the default somehow and don't need a prompt. however, when a 2+ collections exist, require it every time
	if colNameErr != nil {
		return fmt.Errorf("Error setting collection-name: %w", colNameErr)
	}

	db, dbErr := database.ConnectToDB()
	if dbErr != nil {
		return fmt.Errorf("Error connecting to database: %w", dbErr)
	}

	getCol, getColErr := collection.GetCollectionByName(db, colName)
	if getColErr != nil {
		return fmt.Errorf("Error fetching collection: %w", getColErr)
	}

	showPass, passErr := cmd.Flags().GetBool("show-password")
	if passErr != nil {
		return fmt.Errorf("Error showing password: %w", passErr)
	}

	getItem, getErr := login.GetLoginItem(db, itemName, int(getCol.ID))
	if getErr != nil {
		return fmt.Errorf("Error fetching login item: %w", getErr)
	}

	// currently printing out empty struct if collection/item name is given
	// but does not exist in database
	cmd.Printf("\n")
	cmd.Printf("Collection: %v\n", getCol.Name)
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
	colName, colNameErr := cmd.Flags().GetString("collection-name")
	if colName == "" {
		colName = "default"
	}
	if colNameErr != nil {
		return fmt.Errorf("Error setting collection-name: %w", colNameErr)
	}

	itemName, itemErr := cmd.Flags().GetString("item-name")
	if itemErr != nil {
		return fmt.Errorf("Error setting item-name: %w", itemErr)
	}

	db, dbErr := database.ConnectToDB()
	if dbErr != nil {
		return fmt.Errorf("Error connecting to database: %w", dbErr)
	}

	colId, colIdErr := collection.GetCollectionByName(db, colName)
	if colIdErr != nil {
		return fmt.Errorf("Error fetching collection: %w", colIdErr)
	}

	newLoginItem, newLoginErr := login.GetLoginItem(db, itemName, int(colId.ID))
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

	moveCol, moveColErr := cmd.Flags().GetString("move-collection")
	if moveColErr != nil {
		return fmt.Errorf("Error moving login item to new collection: %w", moveColErr)
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

	if len(moveCol) > 0 {
		newColId, newColIdErr := collection.GetCollectionByName(db, moveCol)
		if newColIdErr != nil {
			return fmt.Errorf("Error moving login item to new collection: %w", newColIdErr)
		}
		newLoginItem.CollectionID = int(newColId.ID)
	}

	login.UpdateLoginItem(db, newLoginItem)
	runGetLogin(cmd, args)

	return nil
}

func runListLogins(cmd *cobra.Command, args []string) error {
	colName, colNameErr := cmd.Flags().GetString("collection-name")
	if colNameErr != nil {
		return fmt.Errorf("Error setting collection-name: %w", colNameErr)
	}

	db, dbErr := database.ConnectToDB()
	if dbErr != nil {
		return fmt.Errorf("Error connecting to database: %w", dbErr)
	}

	colId, colIdErr := collection.GetCollectionByName(db, colName)
	if colIdErr != nil {
		return fmt.Errorf("Error fetching collection: %w", colIdErr)
	}

	items, itemErr := login.ListLoginItems(db, int(colId.ID))
	if itemErr != nil {
		return fmt.Errorf("Error fetching login items: %w", itemErr)
	}

	for _, item := range *items {
		cmd.Println(item.ItemName)
	}
	return nil
}

func runDeleteLogin(cmd *cobra.Command, args []string) error {
	colName, colNameErr := cmd.Flags().GetString("collection-name")
	if colName == "" {
		colName = "default"
	}
	if colNameErr != nil {
		return fmt.Errorf("Error setting collection-name: %w", colNameErr)
	}

	db, dbErr := database.ConnectToDB()
	if dbErr != nil {
		return fmt.Errorf("Error connecting to database: %w", dbErr)
	}

	itemToDel, itemDelErr := cmd.Flags().GetString("item-name")
	if itemDelErr != nil {
		return fmt.Errorf("Error finding item to delete: %w", itemDelErr)
	}

	colId, colIdErr := collection.GetCollectionByName(db, colName)
	if colIdErr != nil {
		return fmt.Errorf("Error fetching collection: %w", colIdErr)
	}

	confirm, confirmErr := prompt.ConfirmItemDelete()
	if confirmErr != nil {
		cmd.Printf("Operation cancelled.\n")
		return nil
	}

	if confirm == "y" || confirm == "Y" {
		delErr := login.DeleteLoginItem(db, itemToDel, int(colId.ID))
		if delErr != nil {
			return fmt.Errorf("Error deleting item: %w", delErr)
		}
		cmd.Printf("%v has been deleted.\n", itemToDel)
	}

	return nil
}
