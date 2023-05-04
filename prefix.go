package goarg

type Prefix uint

const (
	PrefixDash Prefix = iota
	PrefixDoubleDash
	PrefixSlash
)

func (p Prefix) String() string {
	switch p {
	case PrefixDash:
		return "-"
	case PrefixDoubleDash:
		return "--"
	case PrefixSlash:
		return "/"
	}

	return "unknown"
}
