// Code generated by "stringer -type=Token"; DO NOT EDIT

package lex

import "fmt"

const _Token_name = "NilDataTokenArrayOpenArrayCloseDictOpenDictClose"

var _Token_index = [...]uint8{0, 3, 12, 21, 31, 39, 48}

func (i Token) String() string {
	if i < 0 || i >= Token(len(_Token_index)-1) {
		return fmt.Sprintf("Token(%d)", i)
	}
	return _Token_name[_Token_index[i]:_Token_index[i+1]]
}
