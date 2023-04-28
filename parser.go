package goarg

type ArgParser struct {
	options []*Option
	values  map[string]string
}

func NewParser() *ArgParser {
	return &ArgParser{
		values: make(map[string]string),
	}
}

func (parser *ArgParser) AddOption(option ...*Option) {
	parser.options = option
}

func (parser *ArgParser) Parse(args []string) error {
	for _, opt := range parser.options {
		val, err := opt.parse(args)

		if err != nil {
			return err
		}

		parser.values[opt.id] = val
	}

	return nil
}

func (parser *ArgParser) Value(id string) (string, bool) {
	val, ok := parser.values[id]
	return val, ok
}
