package cmd

import (
	"flag"
	"github.com/evgbay/pg/common"
)

var (
	// Flag that specifies the password length
	// Example: -l 12
	// By default: 16 symbols
	PasswordLength = flag.Int("l", common.DefaultPasswordLength, "Desired password length")
	// Flag that determines the presence of special characters when creating a password.
	// Example: -s (enable)
	// By default - (disable)
	EnableSpecialSymbols = flag.Bool("s", common.DefaultEnableSpecialSymbols, "Whether to use special characters when creating a password")
)

func ParseArgs() {
	flag.Parse()
}
