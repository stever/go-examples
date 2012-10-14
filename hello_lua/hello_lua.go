package main

/*
#include <stdlib.h>
#include <lua.h>
#include <lualib.h>
#include <lauxlib.h>

// The following functions are wrappers around macros that can't be called from the Lua code.
static void call(lua_State *L, int nargs, int nresults) { return lua_call(L, nargs, nresults); }
static int dofile(lua_State *L, const char *filename) {	return luaL_dofile(L, filename); }
static void getglobal(lua_State *L, const char *name) { return lua_getglobal(L, name); }
static const char *tostring(lua_State *L, int index) { return lua_tostring(L, index); }

#cgo windows LDFLAGS: -llua52
#cgo darwin LDFLAGS: -llua
#cgo linux LDFLAGS: -llua
*/
import "C"
import "unsafe"
import "fmt"

func main() {

	// Initialize Lua.
	var L *C.lua_State = C.luaL_newstate()
	defer C.lua_close(L)

	// Load the Lua libraries.
	C.luaL_openlibs(L)

	// Run the hello.lua script.
	var filename *C.char = C.CString("hello.lua")
	defer C.free(unsafe.Pointer(filename))
	if C.dofile(L, filename) == 0 {

		// Call the 'foo' function.
		var f *C.char = C.CString("foo")
		defer C.free(unsafe.Pointer(f))
		C.getglobal(L, f)
		C.call(L, 0, 0)

		fmt.Println("Called the 'foo' function in", C.GoString(filename))
	} else {
		fmt.Println("Error, unable to run", C.GoString(filename))
		fmt.Println(C.GoString(C.tostring(L, -1)))
	}
}
