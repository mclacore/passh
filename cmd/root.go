/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"strconv"

	"github.com/mclacore/passh/pkg/prompt"
	"github.com/mclacore/passh/pkg/password"
	"github.com/spf13/cobra"
)

var rootCmdLong = `
 ______  ______   ______   ______   __  __    
/\  == \/\  __ \ /\  ___\ /\  ___\ /\ \_\ \   
\ \  _-/\ \  __ \\ \___  \\ \___  \\ \  __ \  
 \ \_\   \ \_\ \_\\/\_____\\/\_____\\ \_\ \_\ 
  \/_/    \/_/\/_/ \/_____/ \/_____/ \/_/\/_/ 
                                              

CLI-based password manager, because why not?
`

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "passh",
	Short: "CLI-based password manager",
	Long:  rootCmdLong,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		persistPass := os.Getenv("PASSH_PERSISTENT_PASS")
		tempPass := os.Getenv("PASSH_PASS")
		timeout, timeoutErr  := strconv.Atoi(os.Getenv("MASTER_PASS_TIMEOUT"))
		if timeoutErr != nil {
			os.Exit(5)
		}

		if persistPass != "" {
			if err := password.ValidateMasterPassword(persistPass); err != nil {
				os.Exit(2)
			}
		} else if tempPass != "" {
			if err := password.ValidateMasterPassword(tempPass); err != nil {
				os.Exit(3)
			}
		} else {
			passInput, passInputErr := prompt.GetMasterPassword()
			if passInputErr != nil {
				os.Exit(4)
			}
			os.Setenv("PASSH_PASS", passInput)
			if os.Getenv("MASTER_PASS_TIMEOUT") == "" {
				os.Setenv("MASTER_PASS_TIMEOUT", "900")
			}
			go password.MasterPasswordTimeout(timeout)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.AddCommand(NewCmdPass())
	rootCmd.AddCommand(NewCmdLogin())
	rootCmd.AddCommand(NewCmdCollection())
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.passh.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
