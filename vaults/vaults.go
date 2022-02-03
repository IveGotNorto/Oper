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
	Items    *items.MapItems
	numItems int
}

type VaultChannel struct {
	Index int
	Items *items.MapItems
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

func (v *Vaults) Display() error {
	for i, vault := range *v {
		for k := range *(*v)[i].Items {
			fmt.Printf("%v\n", vault.Name+"/"+(*(*(*v)[i].Items)[k]).Overview.Title)
		}
	}
	return nil
}

func (v *Vaults) Show(passwordName string) error {
	split := strings.Split(passwordName, "/")
	vaultName := split[0]
	passwordName = split[1]

	for i, vault := range *v {
		if vault.Name != vaultName {
			continue
		}
		if _, ok := (*(*v)[i].Items)[passwordName]; ok {
			out, err := exec.Command("op", "--cache", "get", "item", (*(*(*v)[i].Items)[passwordName]).Uuid, "--fields", "password").Output()
			fmt.Printf("%v", string(out))
			return err
		}
	}
	return nil
}

func PrettyPrint(title string, vaults *Vaults) error {
	fmt.Printf("%v\n", title)
	numVaults := len(*vaults) - 1
	var count int
	for i, vault := range *vaults {
		if i != numVaults {
			fmt.Printf("├── %v\n", vault.Name)
		} else {
			fmt.Printf("└── %v\n", vault.Name)
		}
		count = 0
		for _, item := range *vault.Items {
			if i != numVaults {
				if count != vault.numItems-1 {
					fmt.Printf("│   ├── %v\n", item.Overview.Title)
				} else {
					fmt.Printf("│   └── %v\n", item.Overview.Title)
				}
			} else {
				if count != vault.numItems-1 {
					fmt.Printf("    ├── %v\n", item.Overview.Title)
				} else {
					fmt.Printf("    └── %v\n", item.Overview.Title)
				}
			}
			count++
		}
	}
	return nil
}

func (v *Vaults) Find(sub []string) (Vaults, error) {
	var vaults Vaults
	found := false
	for i, vault := range *v {
		items := make(items.MapItems)
		if contains(vault.Name, sub) {
			vaults = append(vaults, (*v)[i])
		} else {
			for key := range *(*v)[i].Items {
				if contains(key, sub) {
					found = true
					items[key] = (*(*v)[i].Items)[key]
				}
			}
			if found {
				tmp := new(Vault)
				*tmp = (*v)[i]
				tmp.Items = &items
				tmp.numItems = len(items)
				vaults = append(vaults, *tmp)
			}
			found = false
		}
	}
	return vaults, nil
}

func contains(val string, sub []string) bool {
	equal := false
	for _, s := range sub {
		if strings.Contains(val, s) {
			equal = true
			break
		}
	}
	return equal
}
