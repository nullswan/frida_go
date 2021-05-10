package frida_go

/*
 #include "frida-core.h"
*/
import "C"
import (
	"github.com/c3b5aw/go-utils/log"
)

var deviceManager *DeviceManager

func (dm *DeviceManager) init() (err error) {
	manager, err := C.frida_device_manager_new()
	if err != nil {
		log.Error("cannot create device manager")
		return DEVICE_MANAGER_FAIL
	} else {
		log.Info("new device manager")
		dm.ptr = manager
	}
	return
}

func (dm *DeviceManager) GetDeviceByType(dtype C.FridaDeviceType, timeout int) (d *Device, err error) {
	var gerr *C.GError
	cancel := C.g_cancellable_new()
	dev := C.frida_device_manager_get_device_by_type_sync(dm.ptr, dtype, C.gint(timeout), cancel, &gerr)
	if gerr != nil {
		return nil, DEVICE_NOT_FOUND
	}
	return NewDevice(dev)
}

func (d *Device) fromFridaDevice() (err error) {
	d.Name = C.GoString(C.frida_device_get_name(d.ptr))
	d.ID = C.GoString(C.frida_device_get_id(d.ptr))
	d.Type = uint(C.frida_device_get_dtype(d.ptr))
	return
}

func (d *Device) Attach(pid uint) (s *Session, err error) {
	var gerr *C.GError
	cancel := C.g_cancellable_new()
	sess := C.frida_device_attach_sync(d.ptr, C.uint(pid), cancel, &gerr)
	if gerr != nil {
		return nil, UNABLET_TO_ATTACH_DEVICE
	} else {
		s, err = NewSession(d, sess)
	}
	return
}

func NewDevice(fd *C.FridaDevice) (d *Device, err error) {
	d = &Device{ptr: fd}
	err = d.fromFridaDevice()
	return
}

func NewDeviceManager() (dm *DeviceManager, err error) {
	dm = new(DeviceManager)
	err = dm.init()
	return
}

func GetDeviceManager() *DeviceManager {
	if deviceManager == nil {
		deviceManager, _ = NewDeviceManager()
	}
	return deviceManager
}

func GetLocalDevice() (*Device, error) {
	dm := GetDeviceManager()
	return dm.GetDeviceByType(C.FRIDA_DEVICE_TYPE_LOCAL, 10)
}
