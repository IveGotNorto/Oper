package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
	"sort"
)

type URLs struct {
	l string
	u string
}

type Overview struct {
	URLs  []URLs
	Ainfo string
	Title string
	Url   string
}

type Item struct {
	Uuid      string
	Trashed   string
	VaultUuid string
	Overview  Overview
}

type Items []Item

func (i *Items) Len() int           { return len(*i) }
func (i *Items) Less(a, b int) bool { return (*i)[a].Overview.Title < (*i)[b].Overview.Title }
func (i *Items) Swap(a, b int)      { (*i)[a], (*i)[b] = (*i)[b], (*i)[a] }

func (i *Items) retrieve() error {
	out, err := exec.Command("op", "--cache", "list", "items", "--categories", "Login").Output()
	if err != nil {
		return err
	}
	json.Unmarshal(out, &i)
	sort.Sort(i)
	return nil
}

func (i *Items) display() error {
	err := i.retrieve()
	if err != nil {
		return err
	}
	for _, item := range *i {
		fmt.Printf("%v\n", item.Overview.Title)
	}
	return nil
}

func (i *Items) prettyPrint() error {
	return nil
}

func (i *Items) find(pass []string) (Item, error) {
	return Item{}, nil
}

func (i *Items) show(pass string) (string, error) {
	for j := range *i {
		//fmt.Printf("%v ?? %v\n", pass, (*i)[j].Overview.Title)
		if pass == (*i)[j].Overview.Title {
			out, err := exec.Command("op", "--cache", "get", "item", (*i)[j].Uuid, "--fields", "password").Output()
			if err != nil {
				return "", err
			}
			return string(out), nil
		}
	}
	return "", errors.New("password not found")
}
