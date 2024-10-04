package cmd

import (
	"fmt"

	col "github.com/mclacore/passh/pkg/collection"
	"github.com/mclacore/passh/pkg/database"
	"github.com/mclacore/passh/pkg/prompt"
	"github.com/spf13/cobra"
)

func NewCmdCollection() *cobra.Command {
	colCmd := &cobra.Command{
		Use:     "collection",
		Aliases: []string{"col"},
		Short:   "Collection of login items (e.g., for personal, work, school)",
	}

	colNewCmd := &cobra.Command{
		Use:   "new",
		Short: "Create a new login collection",
		RunE:  runNewCollection,
	}
	colNewCmd.MarkPersistentFlagRequired("collection-name")
	colNewCmd.PersistentFlags().StringP("collection-name", "c", "", "Name for the collection")

	colListCmd := &cobra.Command{
		Use:   "list",
		Short: "List login collections",
		RunE:  runListCollections,
	}

	colDelCmd := &cobra.Command{
		Use:     "delete",
		Aliases: []string{"del, remove, rm"},
		Short:   "Delete a collection (and ALL LOGIN ITEMS IN IT)",
		RunE:    runDeleteCollection,
	}

	colCmd.AddCommand(colNewCmd)
	colCmd.AddCommand(colListCmd)
	colCmd.AddCommand(colDelCmd)
	return colCmd
}

func runNewCollection(cmd *cobra.Command, args []string) error {
	colName, nameErr := cmd.Flags().GetString("collection-name")
	if nameErr != nil {
		return fmt.Errorf("Error setting collection-name: %w", nameErr)
	}

	collection := col.Collection{
		ColName: colName,
	}

	db, dbErr := database.ConnectToDB()
	if dbErr != nil {
		return fmt.Errorf("Error connecting to database: %V", dbErr)
	}

	createErr := col.CreateCollection(db, collection)
	if createErr != nil {
		return fmt.Errorf("Error creating new collection: %w", createErr)
	}
	return nil
}

func runListCollections(cmd *cobra.Command, args []string) error {
	db, dbErr := database.ConnectToDB()
	if dbErr != nil {
		return fmt.Errorf("Error connecting to database: %v", dbErr)
	}

	cols, colErr := col.ListCollections(db)
	if colErr != nil {
		return fmt.Errorf("Error fetching collections: %w", colErr)
	}

	for _, col := range *cols {
		cmd.Println(col.ColName)
	}
	return nil
}

func runDeleteCollection(cmd *cobra.Command, args []string) error {
	db, dbErr := database.ConnectToDB()
	if dbErr != nil {
		return fmt.Errorf("Error connecting to database: %v", dbErr)
	}	

	colToDel, colDelErr := cmd.Flags().GetString("collection-name")
	if colDelErr != nil {
		return fmt.Errorf("Error finding collection to delete: %w", colDelErr)
	}

	confirm, confirmErr := prompt.ConfirmCollectionDelete()
	if confirmErr != nil {
		cmd.Printf("Operation cancelled.\n")
		return nil
	}

	if confirm == "y" || confirm == "Y" {
		delErr := col.DeleteCollection(db, colToDel)
		if delErr != nil {
			return fmt.Errorf("Error deleting collection: %w", delErr)
		}
		cmd.Printf("%v has been deleted.\n", colToDel)
	}
	return nil
}