# GoArg
GoArg is a commandline argument parser for go.

## API
Overview over the GoArg API.

### Argument Parser
A new argument parser instance can be created with
```
func NewParser() *ArgumentParser
```

### Argument Matcher
An argument matcher is used to match against a slice of string. An arggument matcher implements the `ArgumentMatcher` interface:
```
Match(args []string) (string, bool)
```

#### String matcher
Used for matching against named arguments like `--foo`. The following prefixes are supported:
- `PrefixDash`: -
- `PrefixDoubleDash`: --
- `PrefixSlash`: /


`StringMatcher` are created by:
```
func NewStringMatcher(p Prefix, keyword string, caseSensitive bool) *StringMatcher
``` 

#### Positional matcher
Used for matching against positional argument at a specific position. The first arguments position is given by `1`.

`PositionalMatcher` are created by:
```
func NewPositionalMatcher(pos uint) *PositionalMatcher 
```

### Options
An `Option` is a aggregator for different matchers which is represented by an identifier and a required field. The `id` is used to retriev parsed values from the argument parser.
```
func NewOption(id string, required bool, matcher ...ArgumentMatcher) *Option
```
Options can be added to the argument parser by calling the follwoing function on an `ArgumentParser`
```
func (parser *ArgumentParser) AddOption(option ...*Option)
```

### Parsing and retrieving values
Parse arguments with given options by calling the following on an `ArgumentParser`, an error is returned if a required argument is not present:
```
func (parser *ArgumentParser) Parse(args []string) error
```

Retrieve parsed values by the option's `id` with:
```
func (parser *ArgumentParser) Value(id string) (string, bool) 
```


## Example
```
parser := goarg.NewParser()

var fooMatcher []goarg.ArgumentMatcher
fooMatcher = append(fooMatcher, goarg.NewStringMatcher(PrefixDash, "f", false))
fooMatcher = append(fooMatcher, goarg.NewStringMatcher(PrefixDoubleDash, "foo", false))

var barMatcher []goarg.ArgumentMatcher
barMatcher = append(barMatcher, goarg.NewStringMatcher(PrefixDash, "b", false))
barMatcher = append(barMatcher, goarg.NewStringMatcher(PrefixDoubleDash, "bar", false))

positionalMatcher := goarg.NewPositionalMatcher(1)

optionPositional := goarg.NewOption("pos", true, positionalMatcher)
optionFoo := goarg.NewOption("foo", false, fooMatcher...)
optionBar := goarg.NewOption("bar", false, barMatcher...)

parser.addOption(optionPositional, optionFoo, optionBar)
```

### Read parsed values
Read the parsed values by using the option `id`
```
fooValue, ok := parser.Value("foo")
fooValue, ok := parser.Value("bar")
positionalValue, ok := parser.Value("pos")
```