# GoArg
GoArg is a commandline argument parser for go.

## Usage

### Create Argument Parser
First create a new argument parser instance
```
parser := goarg.NewParser()
```

### Create Options
Next we create our desired argument options. There are two types of argument options `Named Options` and `Positional Options`. `Named Options` can be created by using `NewOption(id string, dashIdentifier string, doubleDashIdentifier string, required bool) *Option` 
```
optionFoo := goarg.NewOption("foo", "f", "foo", false)
```
For `Positional Arguments` use `func NewPositionalOption(id string, pos int, required bool) *Option`. Note that the postion starts with `1`
```
optionBar := goarg.NewPostionalOption("bar", 1, true)
```

### Add Options
Next we add our options to the parser instance
```
parser.AddOptions(optionFoo, optionBar)
```

### Parse arguments
```
args := os.Args
err := parser.Parse(args)

if err != nil {
    ...
}
```

### Read parsed values
Read the parsed values by using the argument `id`
```
fooValue, ok := parser.Value("foo")
```