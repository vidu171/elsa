package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/divy-work/done/cmd"
	"github.com/lithdew/quickjs"
)

func DoneNS(perms cmd.Perms) func(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
	return func(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
		switch args[0].Int32() {
		case FSRead:
			if !perms.Fs {
				LogError("Perms Error: ", "Filesystem access is blocked.")
				os.Exit(1)
			}
			file := args[1].String()
			dat, e := ioutil.ReadFile(file)
			if e != nil {
				panic(e)
			}
			val := ctx.String(string(dat))
			defer val.Free()
			return val
		case Log:
			fmt.Println(args[1].String())
			return ctx.Null()
		default:
			return ctx.Null()
		}
	}
}