package prompt

import (
	"github.com/manifoldco/promptui"
	"github.com/mclacore/passh/pkg/password"
)

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
		Label: "WARNING: DELETES ALL LOGINS IN COLLECTION. Confirm deletion",
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
		Label: "Enter your Passh master password:",
		IsConfirm: false,
		Mask: rune(2371),
		HideEntered: true,
		Validate: password.ValidateMasterPassword,
	}
	
	result, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return result, nil
}
