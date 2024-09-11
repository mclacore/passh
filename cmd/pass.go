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
		Run:   runNewPass,
	}
	passNewCmd.Flags().Int64P("length", "l", 12, "Length of the password")
	passNewCmd.Flags().BoolP("uppercase", "u", false, "Include uppercase characters")
	passNewCmd.Flags().BoolP("exclude-lowercase", "x", false, "EXCLUDE lowercase characters")
	passNewCmd.Flags().BoolP("numbers", "n", false, "Include numbers")
	passNewCmd.Flags().BoolP("special", "s", false, "Include special characters")

	passCmd.AddCommand(passNewCmd)
	return passCmd
}

func runNewPass(cmd *cobra.Command, args []string) {
	length, _ := cmd.Flags().GetInt64("length")
	uppercase, _ := cmd.Flags().GetBool("uppercase")
	lowercase, _ := cmd.Flags().GetBool("lowercase")
	numbers, _ := cmd.Flags().GetBool("numbers")
	special, _ := cmd.Flags().GetBool("special")

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
}
