package scanner

import (
	"strings"
	"unicode"
)

const (
	IdLen = 32
)
	
type spec struct {
	Id 	string
	Sym int
} 

func asStr(runeArray [IdLen]rune) string {
	n := strings.IndexRune(string(runeArray[:IdLen]), '\x00')
	return string(runeArray[:n])
}

func isEq(s spec, id string, sym int) bool {

	isEqual := false
	
	switch {
		case sym == Symbol || sym == String:
		 	isEqual = (id == s.Id && sym == s.Sym)
		 	
		default:
			isEqual = (sym == s.Sym)
	}

	return isEqual
}

func isNotEq(s spec, id string, sym int) bool {
	return !isEq(s, id, sym)
}

func isNotLetter(ch rune) bool {
	return unicode.ToUpper(ch) < 'A' || 
		   unicode.ToUpper(ch) > 'Z'
}