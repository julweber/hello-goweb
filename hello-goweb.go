package main

import (
    "github.com/hoisie/web"
	"fmt"
)

func hello(ctx *web.Context, val string) string { 

	// retrieve params
    for k,v := range ctx.Params {
        fmt.Println(k, v)
    }
	
	// retrieve route (here: val)
	return "hello " + val
}   

func main() {
    web.Get("/(.*)", hello)
    web.Run("0.0.0.0:9999")
}