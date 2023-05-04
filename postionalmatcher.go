package goarg

type PostionalMatcher struct {
	pos uint
}

func NewPositionalMatcher(pos uint) *PostionalMatcher {
	return &PostionalMatcher{
		pos,
	}
}

func (m PostionalMatcher) Match(args []string) (string, bool) {
	if len(args) > int(m.pos) {
		return args[m.pos], true
	}

	return "", false
}
