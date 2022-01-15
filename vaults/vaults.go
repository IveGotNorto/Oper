package vaults

import (
	"fmt"
	items "oper/items"
	"os/exec"
	"strings"
)

//easyjson:json
type Vault struct {
	Uuid     string `json:"uuid"`
	Name     string `json:"name"`
	Items    *items.Items
	numItems int
}

type VaultChannel struct {
	Index int
	Items *items.Items
	Error error
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
	err = v.retrieveItems()
	if err != nil {
		return err
	}
	return nil
}

func (v *Vaults) retrieveItems() error {
	c := make(chan VaultChannel)
	for i, vault := range *v {
		go func(vs Vault, in int) {
			vItems, err := items.RetrieveByVault(vs.Uuid)
			c <- VaultChannel{in, vItems, err}
		}(vault, i)
	}

	var vc VaultChannel
	for range *v {
		vc = <-c
		(*v)[vc.Index].Items = vc.Items
		(*v)[vc.Index].numItems = len(*vc.Items)
	}
	close(c)
	return nil
}

func (v *Vaults) Display() {
	for i, vault := range *v {
		for _, item := range *(*v)[i].Items {
			fmt.Printf("%v\n", vault.Name+"/"+item.Overview.Title)
		}
	}
}

func (v *Vaults) Show(passwordName string) {
	split := strings.Split(passwordName, "/")
	vaultName := split[0]
	passwordName = split[1]

	for i, vault := range *v {
		if vault.Name != vaultName {
			break
		}
		for j := range *(*v)[i].Items {
			if passwordName == (*(*v)[i].Items)[j].Overview.Title {
				out, _ := exec.Command("op", "--cache", "get", "item", (*(*v)[i].Items)[j].Uuid, "--fields", "password").Output()
				fmt.Printf("%v", string(out))
			}
		}
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
