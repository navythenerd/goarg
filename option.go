package goarg

import (
	"fmt"
)

type Option struct {
	id                   string
	doubleDashIdentifier string
	dashIdentifier       string
	required             bool
	positional           bool
	pos                  int
}

func NewPostionalOption(id string, pos int, required bool) *Option {
	return &Option{
		id:         id,
		required:   required,
		positional: true,
		pos:        pos,
	}
}

func NewOption(id string, dashIdentifier string, doubleDashIdentifier string, required bool) *Option {
	return &Option{
		id:                   id,
		dashIdentifier:       dashIdentifier,
		doubleDashIdentifier: doubleDashIdentifier,
		required:             required,
		positional:           false,
	}
}

func (opt *Option) parse(args []string) (string, error) {
	// check for postional argument
	if opt.positional {

		if (len(args)-1) < opt.pos && opt.required {
			return "", fmt.Errorf("argument %s is required, but was not found", opt.id)
		}

		return args[opt.pos], nil
	}

	// check for dash and double dash identifier values
	for i := 1; i < len(args); i++ {
		if args[i] == opt.dashIdentifier || args[i] == opt.doubleDashIdentifier {
			if len(args) < i+2 {
				return "", fmt.Errorf("identifier for argument %s found, but no value given", opt.id)
			}

			return args[i+1], nil
		}
	}

	// argument required but not found
	if opt.required {
		return "", fmt.Errorf("argument %s is required, but was not found", opt.id)
	}

	// argument is optional
	return "", nil
}
