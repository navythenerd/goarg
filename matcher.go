package goarg

type ArgumentMatcher interface {
	Match(args []string) (string, bool)
}
