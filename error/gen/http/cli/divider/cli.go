// Code generated by goa v3.2.5, DO NOT EDIT.
//
// divider HTTP client CLI support package
//
// Command:
// $ goa gen goa.design/examples/error/design -o
// $(GOPATH)/src/goa.design/examples/error

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	dividerc "goa.design/examples/error/gen/http/divider/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `divider (integer-divide|divide)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` divider integer-divide --a 3601367395041194197 --b 8717444617646084941` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		dividerFlags = flag.NewFlagSet("divider", flag.ContinueOnError)

		dividerIntegerDivideFlags = flag.NewFlagSet("integer-divide", flag.ExitOnError)
		dividerIntegerDivideAFlag = dividerIntegerDivideFlags.String("a", "REQUIRED", "Left operand")
		dividerIntegerDivideBFlag = dividerIntegerDivideFlags.String("b", "REQUIRED", "Right operand")

		dividerDivideFlags = flag.NewFlagSet("divide", flag.ExitOnError)
		dividerDivideAFlag = dividerDivideFlags.String("a", "REQUIRED", "Left operand")
		dividerDivideBFlag = dividerDivideFlags.String("b", "REQUIRED", "Right operand")
	)
	dividerFlags.Usage = dividerUsage
	dividerIntegerDivideFlags.Usage = dividerIntegerDivideUsage
	dividerDivideFlags.Usage = dividerDivideUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "divider":
			svcf = dividerFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "divider":
			switch epn {
			case "integer-divide":
				epf = dividerIntegerDivideFlags

			case "divide":
				epf = dividerDivideFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "divider":
			c := dividerc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "integer-divide":
				endpoint = c.IntegerDivide()
				data, err = dividerc.BuildIntegerDividePayload(*dividerIntegerDivideAFlag, *dividerIntegerDivideBFlag)
			case "divide":
				endpoint = c.Divide()
				data, err = dividerc.BuildDividePayload(*dividerDivideAFlag, *dividerDivideBFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// dividerUsage displays the usage of the divider command and its subcommands.
func dividerUsage() {
	fmt.Fprintf(os.Stderr, `Service is the divider service interface.
Usage:
    %s [globalflags] divider COMMAND [flags]

COMMAND:
    integer-divide: IntegerDivide implements integer_divide.
    divide: Divide implements divide.

Additional help:
    %s divider COMMAND --help
`, os.Args[0], os.Args[0])
}
func dividerIntegerDivideUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] divider integer-divide -a INT -b INT

IntegerDivide implements integer_divide.
    -a INT: Left operand
    -b INT: Right operand

Example:
    `+os.Args[0]+` divider integer-divide --a 3601367395041194197 --b 8717444617646084941
`, os.Args[0])
}

func dividerDivideUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] divider divide -a FLOAT64 -b FLOAT64

Divide implements divide.
    -a FLOAT64: Left operand
    -b FLOAT64: Right operand

Example:
    `+os.Args[0]+` divider divide --a 0.0641408054047874 --b 0.14911855335050145
`, os.Args[0])
}
