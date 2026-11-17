package common

var (
	LowerCase = []rune("abcdefghijklmnopqrstuvwxyz")
	UpperCase = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	Digits    = []rune("0123456789")
	Special   = []rune("@#._&,-(+):*;!?")
	Letters   = append(LowerCase, UpperCase...)
	Alpha     = append(Letters, Digits...)
	All       = append(Alpha, Special...)
)

func GetCharset(template *Template) []rune {
	if template.Special {
		return All
	}
	return Alpha
}
