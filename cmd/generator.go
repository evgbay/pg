package cmd

import (
	"github.com/evgbay/pg/common"
	"math/rand"
	"strings"
)

func GeneratePassword(template *common.Template) string {
	chars := common.GetCharset(template)
	var builder strings.Builder
	builder.Grow(template.Len)
	for i := 0; i < template.Len; i++ {
		builder.WriteRune(chars[rand.Intn(len(chars))])
	}
	return builder.String()
}
