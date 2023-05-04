package goarg

import (
	"fmt"
	"strings"
)

type StringMatcher struct {
	matchStr      string
	caseSensitive bool
}

func NewStringMatcher(p Prefix, keyword string, caseSensitive bool) *StringMatcher {
	return &StringMatcher{
		matchStr:      fmt.Sprintf("%s%s", p.String(), keyword),
		caseSensitive: caseSensitive,
	}
}

func (m *StringMatcher) Match(args []string) (string, bool) {
	for i := 0; i < len(args); i++ {
		str := args[i]

		if !m.caseSensitive {
			str = strings.ToLower(str)
		}

		if str == m.matchStr {
			if len(args) >= i+2 {
				return args[i+1], true
			}

			return "", false
		}
	}

	return "", false
}
