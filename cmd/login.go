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
	loginUpdateCmd.Flags().StringP("move-collection", "m", login.Collection.Name, "Collection to move to")

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
	// need to somehow create a collections table here if a collection is specified but does not exist as a table
	// if no such table of collections is found, it errors out but also automatically drops it in default. BAD
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

	colName, colNameErr := cmd.Flags().GetString("collection-name")
	// remove this part, and instead use promptui to check if a collection exists. if not, prompt user to create one on the spot.
	// if only 1 collection exists, set it as the default somehow and don't need a prompt. however, when a 2+ collections exist, require it every time
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

	getCol, getColErr := collection.GetCollection(db, colName)
	if getColErr != nil {
		return fmt.Errorf("Error fetching collection: %w", getColErr)
	}

	showPass, passErr := cmd.Flags().GetBool("show-password")
	if passErr != nil {
		return fmt.Errorf("Error showing password: %w", passErr)
	}

	itemList, itemListErr := login.ListLoginItems(db, int(getCol.ID))
	if itemListErr != nil {
		return fmt.Errorf("Error fetching login items: %w", itemListErr)
	}

	for _, item := range *itemList {
		if itemName == item.ItemName {
			// GetLoginItem returning the first result twice. this may not be easily possible, and might need to make item_names unique
			getItem, getErr := login.GetLoginItem(db, item.ItemName, int(getCol.ID))
			// OMG SO THIS IS THE PROZBLEM. GET ITEM gets the first on the list. get login either has to return all of the same item name or make itemname unique
			if getItem.ItemName == "" {
				return fmt.Errorf("Cannot find login item named %v in %v collection", itemName, colName)
			}
			if getErr != nil {
				return fmt.Errorf("Error fetching login item: %w", getErr)
			}

			cmd.Printf("---\n")
			cmd.Printf("Collection: %v\n", getCol.Name)
			cmd.Printf("Item Name: %v\n", getItem.ItemName)
			cmd.Printf("Username: %v\n", getItem.Username)

			if showPass {
				cmd.Printf("Password: %v\n", getItem.Password)
			} else {
				cmd.Println("Password: <hidden>")
			}

			cmd.Printf("URL: %v\n", getItem.URL)
		}
	}
	// need to add a for loop here for all listed item-names that match itemName
	// if list that returns login_items.item-name > 1, then for loop here?

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

	colId, colIdErr := collection.GetCollection(db, colName)
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
		newColId, newColIdErr := collection.GetCollection(db, moveCol)
		fmt.Printf("moveCol: %v\n", moveCol)
		fmt.Printf("newColId: %v\n", newColId.ID)
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

	colId, colIdErr := collection.GetCollection(db, colName)
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

	colId, colIdErr := collection.GetCollection(db, colName)
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
