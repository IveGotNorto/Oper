package main

import (
	"oper/app"
	"oper/store"
	"os"
)

func main() {
	var store store.OPStore
	os.Exit(app.Run(store))
}
