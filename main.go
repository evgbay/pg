package main

import (
	"fmt"
	"github.com/evgbay/pg/cmd"
	"github.com/evgbay/pg/common"
)

func main() {
	cmd.ParseArgs()
	template := common.NewTemplate(cmd.PasswordLength, cmd.EnableSpecialSymbols)
	fmt.Println(cmd.GeneratePassword(template))
}
