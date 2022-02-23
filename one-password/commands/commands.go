package commands

import "os/exec"

type OnePasswordCommands struct {
}

func (o *OnePasswordCommands) GetPassword(uniqueId string) (out []byte, err error) {
	out, err = exec.Command("op", "get", "item", uniqueId, "--fields", "password").Output()
	return
}

func (o *OnePasswordCommands) GetContainers() (out []byte, err error) {
	out, err = exec.Command("op", "list", "vaults").Output()
	return
}

func (o *OnePasswordCommands) CreatePasswordInContainer(container string, title string, password string) (err error) {
	_, err = exec.Command("op", "create", "item", "password", "--vault", container, "--title", title, "--tags", "Created in OP", "password=", password).Output()
	return
}

func (o *OnePasswordCommands) CreatePasswordDefaultContainer(title string, password string) (err error) {
	_, err = exec.Command("op", "create", "item", "password", "--title", title, "--tags", "Created in OP", "password=", password).Output()
	return
}

func (o *OnePasswordCommands) GetItemsByContainer(uuid string) (out []byte, err error) {
	out, err = exec.Command("op", "list", "items", "--vault", uuid).Output()
	return
}
