package goarg

import (
	"fmt"
)

type Option struct {
	id       string
	required bool
	matcher  []ArgumentMatcher
}

func NewOption(id string, required bool, matcher ...ArgumentMatcher) *Option {
	return &Option{
		id:       id,
		required: required,
		matcher:  matcher,
	}
}

func (opt *Option) parse(args []string) (string, error) {
	// check arguments against matcher
	for _, m := range opt.matcher {
		val, ok := m.Match(args)

		if ok {
			return val, nil
		}
	}

	// option is required but did not match against any matcher
	if opt.required {
		return "", fmt.Errorf("option %s not found but required", opt.id)
	}

	// did not match, but is not required either
	return "", nil
}
