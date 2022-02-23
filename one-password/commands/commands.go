package commands

import "os/exec"

type OnePasswordCommands struct {
	Cache   bool
	Verbose bool
	Debug   bool
}

func (o *OnePasswordCommands) GetPassword(uniqueId string) (out []byte, err error) {
	if o.Cache {
		out, err = exec.Command("op", "--cache", "get", "item", uniqueId, "--fields", "password").Output()
	} else {
		out, err = exec.Command("op", "get", "item", uniqueId, "--fields", "password").Output()
	}
	return
}

func (o *OnePasswordCommands) GetContainers() (out []byte, err error) {
	if o.Cache {
		out, err = exec.Command("op", "--cache", "list", "vaults").Output()
	} else {
		out, err = exec.Command("op", "list", "vaults").Output()
	}
	return
}

func (o *OnePasswordCommands) CreatePasswordInContainer(container string, title string, password string) (err error) {
	if o.Cache {
		_, err = exec.Command("op", "--cache", "create", "item", "password", "--vault", container, "--title", title, "--tags", "Created in OP", "password=", password).Output()
	} else {
		_, err = exec.Command("op", "create", "item", "password", "--vault", container, "--title", title, "--tags", "Created in OP", "password=", password).Output()
	}
	return
}

func (o *OnePasswordCommands) CreatePasswordDefaultContainer(title string, password string) (err error) {
	if o.Cache {
		_, err = exec.Command("op", "--cache", "create", "item", "password", "--title", title, "--tags", "Created in OP", "password=", password).Output()
	} else {
		_, err = exec.Command("op", "create", "item", "password", "--title", title, "--tags", "Created in OP", "password=", password).Output()
	}
	return
}

func (o *OnePasswordCommands) GetItemsByContainer(uuid string) (out []byte, err error) {
	if o.Cache {
		out, err = exec.Command("op", "--cache", "list", "items", "--vault", uuid).Output()
	} else {
		out, err = exec.Command("op", "list", "items", "--vault", uuid).Output()
	}
	return
}
