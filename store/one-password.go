package store

import (
	"bytes"
	"errors"
	"fmt"
	"oper/one-password/commands"
	"oper/one-password/vaults"
	"strings"
	"syscall"

	"golang.org/x/term"
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
	return pass.Retrieve()
}

func (st OPStore) List() error {
	return pass.List()
}

func (st OPStore) TreeList() error {
	return pass.TreeList()
}

func (st OPStore) Find(terms []string) error {
	entries, err := pass.FindPassword(terms)
	heading := "Search Terms: "
	constructTerms(&heading, terms)
	if err == nil {
		err = vaults.Print(heading, &entries)
	}
	return err
}

func (st OPStore) Show(passName string) error {
	buf, err := pass.Show(passName)
	fmt.Println(buf)
	return err
}

func (st OPStore) Insert(vaultPass string) error {
	var vault string
	var passName string
	// Split on first occurrence of /
	tmp := strings.SplitN(vaultPass, "/", 2)
	if len(tmp) > 1 {
		vault = tmp[0]
		passName = tmp[1]
		// verify vault
		if !pass.VerifyContainerByName(vault) {
			return errors.New("vault '" + vault + "' not found")
		}
	}

	fmt.Printf("Enter password for %v: ", vaultPass)
	password, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return err
	}
	fmt.Printf("\nRetype password for %v: ", vaultPass)
	buff, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return err
	}

	if !bytes.Equal(buff, password) {
		return errors.New("the entered passwords do not match")
	}

	return pass.Insert(vault, passName, string(password))
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

func (st OPStore) Edit() error {
	return nil
}

func (st OPStore) Generate() error {
	return nil
}

func (st OPStore) Remove() error {
	return nil
}

func (st OPStore) Move() error {
	return nil
}

func (st OPStore) Copy() error {
	return nil
}
