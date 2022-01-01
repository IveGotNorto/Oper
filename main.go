package main

import (
	items "oper/items"
)

type OPDisplay interface {
	Display()
	PrettyPrint() error
	Retrieve() error
	Find([]string) (items.Item, error)
}

type OPInteract interface {
	Show(string) error
}

func OpPrint(op OPDisplay) {
	op.Display()
}

func OpPrettyPrint(op OPDisplay) error {
	err := op.PrettyPrint()
	return err
}

func OpFind(op OPDisplay, pass []string) error {
	var err error
	return err
}

func OpShow(op OPInteract, pass string) error {
	err := op.Show(pass)
	return err
}
