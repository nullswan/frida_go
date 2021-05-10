package frida_go

/*
 #include "frida-core.h"
*/
import "C"
import "unsafe"

const (
	SCRIPT_RUNTIME_V8 = uint(C.FRIDA_SCRIPT_RUNTIME_V8)
)

func NewScript(sess *Session, name string, source interface{}, runtime C.FridaScriptRuntime) (s *Script, err error) {
	opts := C.frida_script_options_new()
	defer func() {
		C.frida_unref(C.gpointer(opts))
		opts = nil
	}()

	C.frida_script_options_set_name(opts, C.CString(name))
	C.frida_script_options_set_runtime(opts, runtime)

	var (
		gerr   *C.GError
		script *C.FridaScript
	)
	cancel := C.g_cancellable_new()
	switch src := source.(type) {
	case string:
		script = C.frida_session_create_script_sync(sess.ptr, C.CString(src), opts, cancel, &gerr)
	case []byte:
		if gBytes, ok := GoBytesToGBytes(src); ok {
			script = C.frida_session_create_script_from_bytes_sync(sess.ptr, gBytes, opts, cancel, &gerr)
		}
	}
	if gerr != nil {
		return nil, CREATE_SCRIPT_ERROR
	} else if !IsNullCPointer(unsafe.Pointer(script)) {
		s = &Script{
			ptr:  script,
			ID:   uint(C.frida_script_get_id(script)),
			Name: name,
		}
	}
	return
}

func (scr *Script) Load() error {
	var gerr *C.GError
	cancel := C.g_cancellable_new()
	C.frida_script_load_sync(scr.ptr, cancel, &gerr)
	if gerr != nil {
		return LOAD_SCRIPT_ERROR
	}
	return nil
}
