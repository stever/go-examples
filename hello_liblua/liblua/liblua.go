package lua

/*
#cgo pkg-config: lua51

#include <stdlib.h>
#include <lua.h>
#include <lualib.h>
#include <lauxlib.h>

// The following functions are wrappers around macros that can't be called from the Lua code.
static void call(lua_State *L, int nargs, int nresults) { return lua_call(L, nargs, nresults); }
static int dofile(lua_State *L, const char *filename) {	return luaL_dofile(L, filename); }
static void getglobal(lua_State *L, const char *name) { return lua_getglobal(L, name); }
static const char *tostring(lua_State *L, int index) { return lua_tostring(L, index); }
*/
import "C"
import "unsafe"

type Context struct {
	lua *C.lua_State
}

func NewContext() *Context {
	var ctx Context
	ctx.Init()
	return &ctx
}

func (ctx *Context) Init() {
	ctx.lua = C.luaL_newstate()
	C.luaL_openlibs(ctx.lua)
}

func (ctx *Context) Close() {
	C.lua_close(ctx.lua)
}

func (ctx *Context) Do(filename string) (ok bool, err string) {
	var c_filename *C.char = C.CString(filename)
	defer C.free(unsafe.Pointer(c_filename))
	ok = C.dofile(ctx.lua, c_filename) == 0
	err = C.GoString(C.tostring(ctx.lua, -1))
	return
}

func (ctx *Context) Call(name string) {
	var c_name *C.char = C.CString("foo")
	defer C.free(unsafe.Pointer(c_name))
	C.getglobal(ctx.lua, c_name)
	C.call(ctx.lua, 0, 0)
}
