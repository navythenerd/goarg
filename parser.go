package goarg

type ArgumentParser struct {
	options []*Option
	values  map[string]string
}

func NewParser() *ArgumentParser {
	return &ArgumentParser{
		values: make(map[string]string),
	}
}

func (parser *ArgumentParser) AddOption(option ...*Option) {
	parser.options = option
}

func (parser *ArgumentParser) Parse(args []string) error {
	for _, opt := range parser.options {
		val, err := opt.parse(args)

		if err != nil {
			return err
		}

		parser.values[opt.id] = val
	}

	return nil
}

func (parser *ArgumentParser) Value(id string) (string, bool) {
	val, ok := parser.values[id]
	return val, ok
}
