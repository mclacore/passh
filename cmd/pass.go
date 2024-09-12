package cmd

import (
	"math/rand"

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
	passNewCmd.Flags().BoolP("uppercase", "u", false, "Include uppercase characters")
	passNewCmd.Flags().BoolP("exclude-lowercase", "x", false, "EXCLUDE lowercase characters")
	passNewCmd.Flags().BoolP("numbers", "n", false, "Include numbers")
	passNewCmd.Flags().BoolP("special", "s", false, "Include special characters")

	passCmd.AddCommand(passNewCmd)
	return passCmd
}

func runNewPass(cmd *cobra.Command, args []string) error {
	length, lenErr := cmd.Flags().GetInt64("length")
	if lenErr != nil {
		return lenErr
	}

	uppercase, upperErr := cmd.Flags().GetBool("uppercase")
	if upperErr != nil {
		return upperErr
	}

	lowercase, lowerErr := cmd.Flags().GetBool("lowercase")
	if lowerErr != nil {
		return lowerErr
	}

	numbers, numErr := cmd.Flags().GetBool("numbers")
	if numErr != nil {
		return numErr
	}

	special, specErr := cmd.Flags().GetBool("special")
	if specErr != nil {
		return specErr
	}

	charset := ""
	if uppercase {
		charset += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if !lowercase {
		charset += "abcdefghijklmnopqrstuvwxyz"
	}
	if numbers {
		charset += "0123456789"
	}
	if special {
		charset += "!@#$%^&*()_+~><"
	}

	password := make([]byte, length)
	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}

	cmd.Println(string(password))

	return nil
}
