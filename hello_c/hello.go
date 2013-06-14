package hello

/*
#include "hello.h"
*/
import "C"

var Greeting = C.GoString(C.greeting)
