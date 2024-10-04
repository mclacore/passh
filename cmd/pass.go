package cmd

import (
	"fmt"

	"github.com/mclacore/passh/pkg/password"
	"github.com/spf13/cobra"
)

func NewCmdPass() *cobra.Command {
	passCmd := &cobra.Command{
		Use:   "pass",
		Short: "Generate, store, and retrieve passwords",
	}

	passNewCmd := &cobra.Command{
		Use:   "new",
		Short: "Generate a new password",
		RunE:  runNewPass,
	}
	passNewCmd.Flags().Int64P("length", "l", 12, "Length of the password")
	passNewCmd.Flags().BoolP("uppercase", "u", false, "Include uppercase characters") // should be set to true
	passNewCmd.Flags().BoolP("exclude-lowercase", "x", false, "EXCLUDE lowercase characters")
	passNewCmd.Flags().BoolP("numbers", "n", false, "Include numbers") // should be set to true
	passNewCmd.Flags().BoolP("special", "s", false, "Include special characters") // should be set to true

	passCmd.AddCommand(passNewCmd)
	return passCmd
}

func runNewPass(cmd *cobra.Command, args []string) error {
	length, lenErr := cmd.Flags().GetInt64("length")
	if lenErr != nil {
		return fmt.Errorf("Error setting password length: %v", lenErr)
	}

	uppercase, upperErr := cmd.Flags().GetBool("uppercase")
	if upperErr != nil {
		return fmt.Errorf("Error enabling uppercase chars: %v", upperErr)
	}

	lowercase, lowerErr := cmd.Flags().GetBool("exclude-lowercase")
	if lowerErr != nil {
		return fmt.Errorf("Error excluding lowercase chars: %v", lowerErr)
	}

	numbers, numErr := cmd.Flags().GetBool("numbers")
	if numErr != nil {
		return fmt.Errorf("Error enabling number chars: %v", numErr)
	}

	special, specErr := cmd.Flags().GetBool("special")
	if specErr != nil {
		return fmt.Errorf("Error enabling special chars: %v", specErr)
	}

	password := password.GeneratePassword(length, uppercase, lowercase, numbers, special)

	cmd.Println(password)

	return nil
}
