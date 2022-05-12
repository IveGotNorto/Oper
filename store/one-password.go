package store

import (
	"fmt"
	"oper/one-password/commands"
	"oper/one-password/vaults"
	"strings"
)

var pass vaults.Vaults

type OPStore struct {
}

func (st OPStore) Setup(args StoreArguments) error {
	vaults.Commands = commands.OnePasswordCommands{
		Cache:   args.Cache,
		Verbose: args.Verbose,
		Debug:   args.Debug,
	}
	err := pass.Retrieve()
	if err != nil {
		return fmt.Errorf("unable to retrieve one password items")
	}
	return err
}

func (st OPStore) List(order string) error {
	return pass.List(selectSortingOrder(order))
}

func (st OPStore) TreeList(order string) error {
	return pass.TreeList(selectSortingOrder(order))
}

func (st OPStore) Show(passName string) error {
	buf, err := pass.Show(passName)
	fmt.Println(buf)
	return err
}

func selectSortingOrder(order string) int {
	var enumOrder int
	if strings.Contains(order, "desc") {
		enumOrder = Descending
	} else {
		enumOrder = Ascending
	}
	return enumOrder
}
