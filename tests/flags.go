package tests

import (
	"gopkg.in/urfave/cli.v1"
)

var (
	gasLimit = cli.IntFlag{
		Name:  "gaslimit",
		Usage: "gaslimit ",
		Value: 10000000,
	}
)

var (
	GasLimit = uint64(10000000)
)

// Flags holds all command-line flags required for tests.
var Flags = []cli.Flag{
	gasLimit,
}

func SetParams(ctx *cli.Context) {

	if ctx.GlobalIsSet(gasLimit.Name) {
		GasLimit = ctx.GlobalUint64(gasLimit.Name)
	}

}
