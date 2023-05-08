package goarg

type PositionalMatcher struct {
	pos uint
}

func NewPositionalMatcher(pos uint) *PositionalMatcher {
	return &PositionalMatcher{
		pos,
	}
}

func (m PositionalMatcher) Match(args []string) (string, bool) {
	if len(args) > int(m.pos) {
		return args[m.pos], true
	}

	return "", false
}
