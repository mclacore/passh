/*
Copyright © 2024 Michael Lacore mclacore@gmail.com
*/
package cmd

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mclacore/passh/pkg/database"
	"github.com/mclacore/passh/pkg/env"
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

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "passh",
	Short: "CLI-based password manager",
	Long:  rootCmdLong,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		_ = godotenv.Load(".env")

		if os.Getenv("PASSH_USER") == "" {
			env.SetPasshUserEnv("postgres")
			prompt.WelcomeWizard()
		}

		// Set this if you don't want to re-auth into Passh after timeout
		persistPass := os.Getenv("PASSH_PERSISTENT_PASS")
		tempPass := os.Getenv("PASSH_PASS")

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

			env.SetPasshTempPassEnv(passInput)
			if _, err := database.ConnectToDB(); err != nil {
				log.Print("Invalid password")
				os.Exit(401)
			}

			timeout := os.Getenv("PASSH_TIMEOUT")
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
