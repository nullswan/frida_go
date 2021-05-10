package frida_go

/*
 #cgo CFLAGS: -g -O2 -w -I. -I${SRCDIR}/libs
 #cgo LDFLAGS: -static-libgcc -L${SRCDIR}/libs -lfrida-core -ldl -lm -lrt -lresolv -lpthread -Wl,--export-dynamic
 #include "frida-core.h"
*/
import "C"
import (
	"unsafe"

	"github.com/c3b5aw/go-utils/log"
)

func init() {
	_, err := C.frida_init()
	if err != nil {
		log.Error("cannot start hook system")
	}
	log.Info("hook system started")
}

func IsNullCPointer(ptr unsafe.Pointer) bool {
	return uintptr(ptr) == uintptr(0)
}

func GoBytesToGBytes(bytes []byte) (g *C.GBytes, ok bool) {
	size := len(bytes)
	g = C.g_bytes_new_take(C.gpointer(C.CBytes(bytes)), C.ulong(size))
	if IsNullCPointer(unsafe.Pointer(g)) {
		return nil, false
	}
	gSize := C.g_bytes_get_size(g)
	if int(gSize) != size {
		return nil, false
	}

	return g, true
}
