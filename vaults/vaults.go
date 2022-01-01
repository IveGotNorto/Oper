package vaults

import (
	"fmt"
	items "oper/items"
	"os/exec"
)

//easyjson:json
type Vault struct {
	Uuid     string `json:"uuid"`
	Name     string `json:"name"`
	Items    *items.Items
	numItems int
}

//easyjson:json
type Vaults []Vault

func (v *Vaults) Retrieve() error {
	out, err := exec.Command("op", "--cache", "list", "vaults").Output()
	if err != nil {
		return err
	}
	err = v.UnmarshalJSON(out)
	if err != nil {
		return err
	}
	return nil
}

func (v *Vaults) retrieveItems() error {
	for i, vault := range *v {
		vItems, err := items.RetrieveByVault(vault.Uuid)
		if err != nil {
			return err
		}
		(*v)[i].Items = vItems
		(*v)[i].numItems = len(*vItems)
	}
	return nil
}

func (v *Vaults) Display() {
	for _, vault := range *v {
		fmt.Printf("%v\n", vault.Name)
	}
}

func (v *Vaults) PrettyPrint() error {
	fmt.Printf("One Password Store\n")
	err := v.retrieveItems()
	if err != nil {
		return err
	}
	numVaults := len(*v) - 1

	for i, vault := range *v {

		if i != numVaults {
			fmt.Printf("├── %v\n", vault.Name)
		} else {
			fmt.Printf("└── %v\n", vault.Name)
		}

		for j, item := range *vault.Items {
			if i != numVaults {
				if j != vault.numItems-1 {
					fmt.Printf("│   ├── %v\n", item.Overview.Title)
				} else {
					fmt.Printf("│   └── %v\n", item.Overview.Title)
				}
			} else {
				if j != vault.numItems-1 {
					fmt.Printf("    ├── %v\n", item.Overview.Title)
				} else {
					fmt.Printf("    └── %v\n", item.Overview.Title)
				}
			}
		}
	}
	return nil
}

func (v *Vaults) Find(pass []string) (items.Item, error) {
	return items.Item{}, nil
}
