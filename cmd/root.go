/*
Copyright © 2024 Michael Lacore mclacore@gmail.com
*/
package cmd

import (
	"log"
	"os"

	"github.com/mclacore/passh/pkg/config"
	"github.com/mclacore/passh/pkg/database"
	"github.com/mclacore/passh/pkg/password"
	"github.com/mclacore/passh/pkg/prompt"
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

var rootCmd = &cobra.Command{
	Use:               "passh",
	Short:             "CLI-based password manager",
	Long:              rootCmdLong,
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		user, userErr := config.LoadConfigValue("auth", "username")
		if userErr != nil {
			log.Printf("Error loading username: %v", userErr)
			os.Exit(5)
		}

		if user == "" {
			config.SaveConfigValue("auth", "username", "postgres")
			config.SaveConfigValue("auth", "timeout", "900")
			prompt.WelcomeWizard()
		}

		// Set PASSH_PERSISTENT_PASS if you don't want to keep re-authing after timeout
		persistPass, persistPassErr := config.LoadConfigValue("auth", "persist_pass")
		if persistPassErr != nil {
			log.Printf("Error loading persistent pass: %v", persistPassErr)
			os.Exit(5)
		}
		tempPass, tempPassErr := config.LoadConfigValue("auth", "temp_pass")
		if tempPassErr != nil {
			log.Printf("Error loading temp pass: %v", tempPassErr)
			os.Exit(5)
		}

		if persistPass != "" {
			if _, err := database.ConnectToDB(); err != nil {
				log.Print("Invalid persistent password")
				os.Exit(401)
			}
		} else if tempPass != "" {
			if _, err := database.ConnectToDB(); err != nil {
				log.Print("Invalid password")
				os.Exit(401)
			}
		} else {
			passInput, passInputErr := prompt.GetMasterPassword()
			if passInputErr != nil {
				log.Printf("Something went wrong with inputting password: %v", passInputErr)
				os.Exit(3)
			}

			config.SaveConfigValue("auth", "temp_pass", passInput)
			if _, err := database.ConnectToDB(); err != nil {
				log.Print("Invalid password")
				os.Exit(401)
			}

			timeout, timeoutErr := config.LoadConfigValue("auth", "timeout")
			if timeoutErr != nil {
				log.Printf("Error loading timeout value: %v", timeoutErr)
				os.Exit(5)
			}

			go password.MasterPasswordTimeout(timeout)
		}
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

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
