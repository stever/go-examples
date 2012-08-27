package main

import "../liblua"
import "fmt"

func main() {
	var lua lua.Context
	lua.Init()
	defer lua.Close()
	filename := "hello.lua"
	ok, err := lua.Do(filename)
	if ok {
		lua.Call("foo")
		fmt.Println("Called the 'foo' function in", filename)
	} else {
		fmt.Println("Error, unable to run", filename)
		fmt.Println(err)
	}
}
