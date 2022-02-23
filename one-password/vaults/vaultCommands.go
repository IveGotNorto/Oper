package vaults

import "oper/one-password/commands"

var op commands.OnePasswordCommands

func getPassword(itemUuid string) (string, error) {
	buff, err := op.GetPassword(itemUuid)
	if err != nil {
		return "", err
	}
	return string(buff), err
}

func (v *Vaults) getVaults() error {
	out, err := op.GetContainers()
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
	return op.CreatePasswordDefaultContainer(title, password)
}

func createPasswordInVault(container string, title string, password string) error {
	return op.CreatePasswordInContainer(container, title, password)
}
