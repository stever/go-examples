package lua

/*
#include <stdlib.h>
#include <lua.h>
#include <lualib.h>
#include <lauxlib.h>

static int dofile(lua_State *L, const char *filename) {
	return luaL_dofile(L, filename);
}

static void getglobal(lua_State *L, const char *name) {
	return lua_getglobal(L, name);
}

static const char *tostring(lua_State *L, int index) {
	return lua_tostring(L, index);
}

#cgo windows,386 CFLAGS: -IC:/Lua/include
#cgo windows,386 LDFLAGS: -LC:/Lua/lib
#cgo windows,amd64 CFLAGS: -IC:/Lua_x86_64/include
#cgo windows,amd64 LDFLAGS: -LC:/Lua_x86_64
#cgo windows LDFLAGS: -llua51
#cgo darwin LDFLAGS: -llua
#cgo linux LDFLAGS: -llua
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
	C.lua_call(ctx.lua, 0, 0)
}
