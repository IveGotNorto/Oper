package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type Vault struct {
	Uuid  string
	Name  string
	Items *[]Item
}

type Vaults []Vault

func (v *Vaults) retrieve() error {
	out, err := exec.Command("op", "--cache", "list", "vaults").Output()
	json.Unmarshal(out, &v)
	if err != nil {
		return err
	}

	for i, vault := range *v {
		com, err := exec.Command("op", "--cache", "list", "items", "--vault", vault.Uuid).Output()
		if err != nil {
			return err
		}
		json.Unmarshal(com, &(*v)[i].Items)
	}
	return nil
}

func (v *Vaults) display() error {
	err := v.retrieve()
	if err != nil {
		return err
	}
	for _, vault := range *v {
		fmt.Printf("%v\n", vault.Name)
	}
	return nil
}

func (v *Vaults) prettyPrint() error {
	fmt.Printf("One Password Store\n")

	numVaults := len(*v) - 1

	for i, vault := range *v {

		if i != numVaults {
			fmt.Printf("├── %v\n", vault.Name)
		} else {
			fmt.Printf("└── %v\n", vault.Name)
		}

		numItems := len(*vault.Items) - 1

		for j, item := range *vault.Items {
			if i != numVaults {
				if j != numItems {
					fmt.Printf("│   ├── %v\n", item.Overview.Title)
				} else {
					fmt.Printf("│   └── %v\n", item.Overview.Title)
				}
			} else {
				if j != numItems {
					fmt.Printf("    ├── %v\n", item.Overview.Title)
				} else {
					fmt.Printf("    └── %v\n", item.Overview.Title)
				}
			}
		}
	}

	return nil

}

func (v *Vaults) find(pass []string) (Item, error) {
	return Item{}, nil
}
