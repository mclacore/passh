package prompt

import (
	"errors"

	"github.com/manifoldco/promptui"
)

var promptsWizard = []func() (string, error){
	getHost,
	getUser,
	getPass,
	getPort,
}

var validate = func(input string) error {
	if len(input) < 12 {
		return errors.New("Password must have 12 or more characters.")
	}
	return nil
}

func ConfirmItemDelete() (string, error) {
	prompt := promptui.Prompt{
		Label:     "Confirm deletion of the item",
		IsConfirm: true,
	}

	result, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return result, nil
}

func ConfirmCollectionDelete() (string, error) {
	prompt := promptui.Prompt{
		Label:     "WARNING: DELETES ALL LOGINS IN COLLECTION. Confirm deletion",
		IsConfirm: true,
	}

	result, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return result, nil
}

func GetMasterPassword() (string, error) {
	prompt := promptui.Prompt{
		Label:     "Enter your Passh master password:",
		IsConfirm: false,
		Mask:      '*',
		Validate:  validate,
	}

	result, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return result, nil
}

func WelcomeWizard() error {
	for _, prompt := range promptsWizard {
	}

	return nil
}

func getHost() (string, error) {
	prompt := promptui.Prompt{
		Label:   "Set a database hostname",
		Default: "localhost",
	}

	result, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return result, nil
}

func getUser() (string, error) {
	prompt := promptui.Prompt{
		Label:   "Create a Passh Username",
		Default: "passh",
	}

	result, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return result, nil
}

func getPass() (string, error) {
	prompt := promptui.Prompt{
		Label:    "Create a Passh Password",
		Validate: validate,
		Mask:     '*',
	}

	result, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return result, nil
}

func getPort() (string, error) {
	prompt := promptui.Prompt{
		Label: "Set a database port",
		Default: "5432",
	}

	result, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return result, nil
}

