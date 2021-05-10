package frida_go

/*
 #include "frida-core.h"
*/
import "C"

func (p *Process) fromFridaProcess() (err error) {
	p.Name = C.GoString(C.frida_process_get_name(p.ptr))
	p.Pid = uint(C.frida_process_get_pid(p.ptr))
	return
}

func NewProcess(fp *C.FridaProcess) (p *Process, err error) {
	p = &Process{ptr: fp}
	err = p.fromFridaProcess()
	return
}
