package prompt

import (
	"errors"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/mclacore/passh/pkg/database"
)

var promptsWizard = []func() error{
	getPass,
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
		if err := prompt(); err != nil {
			return err
		}
	}

	return nil
}

func getPass() error {
	prompt := promptui.Prompt{
		Label:    "Create a Passh Password",
		Validate: validate,
		Mask:     '*',
	}

	result, err := prompt.Run()
	if err != nil {
		return err
	}

	database.WizardPasswordSet(result)
	os.Setenv("PASSH_PASS", result)

	return nil
}
