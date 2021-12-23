package main

import (
	"flag"
	"fmt"
	"os"
)

type OPDisplay interface {
	display() error
	prettyPrint() error
	retrieve() error
	find([]string) (Item, error)
}

type OPInteract interface {
	show(string) (string, error)
}

func main() {
	//lsCmd := flag.NewFlagSet("ls", flag.ExitOnError)
	//findCmd := flag.NewFlagSet("find", flag.ExitOnError)
	showCmd := flag.NewFlagSet("show", flag.ExitOnError)
	passFlag := showCmd.String("password", "", "Password to show")

	var items Items
	var vaults Vaults

  if len(os.Args) > 1 {
    switch os.Args[1] {
    // Cases for vaults+items
    case "ls":
      vaults.retrieve()
      opPrettyPrint(&vaults)
    default:
      // Cases for items
      items.retrieve()
      switch os.Args[1] {
      case "find":
        //opFind(&items, "woo")
      case "show":
        showCmd.Parse(os.Args[2:])
        opShow(&items, *passFlag)
      }
    }
  } else {
    items.retrieve()
    opPrint(&items)
  }
}

func opPrint(op OPDisplay) {
	if err := op.display(); err != nil {
		fmt.Println(err)
	}
}

func opPrettyPrint(op OPDisplay) {
	if err := op.prettyPrint(); err != nil {
		fmt.Println(err)
	}
}

func opFind(op OPDisplay, pass []string) {

}

func opShow(op OPInteract, pass string) {
	out, err := op.show(pass)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(out)
}
