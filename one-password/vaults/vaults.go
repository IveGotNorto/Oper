package vaults

import (
	"errors"
	"fmt"
	items "oper/one-password/items"
	"sort"
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
	err := v.getVaults()
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
			if err != nil {
				c <- VaultChannel{in, nil, err}
			} else {
				c <- VaultChannel{in, vItems, nil}
			}
		}(vault, i)
	}

	var vc VaultChannel
	for range *v {
		vc = <-c
		if vc.Error != nil {
			return vc.Error
		}
		(*v)[vc.Index].Items = vc.Items
		(*v)[vc.Index].numItems = len(*vc.Items)
	}
	close(c)
	return nil
}

func (v *Vaults) Sort(order int) {
	asc := func(i, j int) bool { return (*v)[i].Name < (*v)[j].Name }
	desc := func(i, j int) bool { return (*v)[i].Name > (*v)[j].Name }
	if order == 0 {
		sort.Slice(*v, asc)
	} else {
		sort.Slice(*v, desc)
	}
}

func (v *Vaults) List(order int) error {
	v.Sort(order)

	var buff []string
	for i, vault := range *v {
		for key := range *(*v)[i].Items {
			buff = append(buff, key)
		}

		if order == 0 {
			sort.Strings(buff)
		} else {
			sort.Sort(sort.Reverse(sort.StringSlice(buff)))
		}

		for _, keyName := range buff {
			fmt.Printf("%v\n", vault.Name+"/"+keyName)
		}
		buff = nil
	}
	return nil
}

func (v *Vaults) Show(passwordName string) (string, error) {
	split := strings.Split(passwordName, "/")
	vaultName := split[0]
	passwordName = split[1]

	for i, vault := range *v {
		if vault.Name != vaultName {
			continue
		}
		if _, ok := (*(*v)[i].Items)[passwordName]; ok {
			pass, err := getPassword((*(*(*v)[i].Items)[passwordName]).Uuid)
			return pass, err
		}
	}
	return "", errors.New("password not found")
}

func (v *Vaults) TreeList(order int) error {
	v.Sort(order)
	return Print("One Password Store", v, order)
}

func Print(title string, vaults *Vaults, order int) error {
	fmt.Printf("%v\n", title)
	numVaults := len(*vaults) - 1
	var count int

	var keys []string
	for i, vault := range *vaults {
		if i != numVaults {
			fmt.Printf("├── %v\n", vault.Name)
		} else {
			fmt.Printf("└── %v\n", vault.Name)
		}

		// Im sorry
		for k := range *vault.Items {
			keys = append(keys, k)
		}

		if order == 0 {
			sort.Strings(keys)
		} else {
			sort.Sort(sort.Reverse(sort.StringSlice(keys)))
		}

		count = 0
		for _, key := range keys {
			if i != numVaults {
				if count != vault.numItems-1 {
					fmt.Printf("│   ├── %v\n", key)
				} else {
					fmt.Printf("│   └── %v\n", key)
				}
			} else {
				if count != vault.numItems-1 {
					fmt.Printf("    ├── %v\n", key)
				} else {
					fmt.Printf("    └── %v\n", key)
				}
			}
			count++
		}

		// clear keys buffer
		keys = nil
	}
	return nil
}

func (v *Vaults) FindPassword(sub []string) (Vaults, error) {
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

func (v *Vaults) Insert(vault string, title string, password string) error {
	var err error
	password = "password=" + password
	if vault != "" {
		err = createPasswordInVault(vault, title, password)
	} else {
		err = createPassword(title, password)
	}
	return err
}

func (v *Vaults) VerifyContainerByName(name string) bool {
	for _, vault := range *v {
		if vault.Name == name {
			return true
		}
	}
	return false
}
