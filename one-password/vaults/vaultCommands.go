package vaults

import "oper/one-password/commands"

var Commands commands.OnePasswordCommands

func getPassword(itemUuid string) (string, error) {
	buff, err := Commands.GetPassword(itemUuid)
	if err != nil {
		return "", err
	}
	return string(buff), err
}

func (v *Vaults) getVaults() error {
	out, err := Commands.GetContainers()
	if err != nil {
		return err
	}
	err = v.UnmarshalJSON(out)
	if err != nil {
		return err
	}
	return err
}

func createPassword(title string, password string) error {
	return Commands.CreatePasswordDefaultContainer(title, password)
}

func createPasswordInVault(container string, title string, password string) error {
	return Commands.CreatePasswordInContainer(container, title, password)
}
