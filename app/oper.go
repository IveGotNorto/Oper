package app

import (
	items "oper/items"
	"oper/vaults"
)

var v vaults.Vaults

type OPDisplay interface {
	Display() error
	PrettyPrint() error
	Retrieve() error
	Find([]string) (items.Item, error)
}

type OPInteract interface {
	Show(string) error
}

func Setup() error {
	return v.Retrieve()
}

func OpPrint() error {
	return v.Display()
}

func OpPrettyPrint() error {
	return vaults.PrettyPrint("One Password Store", &v)
}

func OpFind(terms []string) error {
	entries, err := v.Find(terms)
	heading := "Search Terms: "
	constructTerms(&heading, terms)
	if err == nil {
		err = vaults.PrettyPrint(heading, &entries)
	}
	return err
}

func OpShow(pass string) error {
	return v.Show(pass)
}

func constructTerms(pre *string, terms []string) {
	length := len(terms) - 1
	for i, term := range terms {
		*pre += term
		if i < length {
			*pre += ","
		}
	}
}
