package frida_go

/*
 #include "frida-core.h"
*/
import "C"

func NewSession(dev *Device, fs *C.FridaSession) (s *Session, err error) {
	s = &Session{
		ptr: fs,
		Dev: dev,
		Pid: uint(C.frida_session_get_pid(fs)),
	}
	return
}

func (sess *Session) CreateScriptSync(name string, source string, runtime ...uint) (s *Script, err error) {
	return NewScript(sess, name, source, C.FRIDA_SCRIPT_RUNTIME_V8)
}
