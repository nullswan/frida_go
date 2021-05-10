package frida_go

/*
 #include "frida-core.h"
*/
import "C"

type Device struct {
	ptr  *C.FridaDevice
	Name string
	ID   string
	Type uint
}

type Session struct {
	ptr *C.FridaSession
	Dev *Device
	Pid uint
}

type DeviceManager struct {
	ptr *C.FridaDeviceManager
}

type Process struct {
	ptr  *C.FridaProcess
	Name string
	Pid  uint
}

type Script struct {
	ptr  *C.FridaScript
	ID   uint
	Name string
}
