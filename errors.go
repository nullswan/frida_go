package frida_go

import "errors"

var (
	ATTACH_SCRIPT_ERROR      = errors.New("unable to attach script")
	DEVICE_MANAGER_FAIL      = errors.New("unable to process device manager")
	DEVICE_NOT_FOUND         = errors.New("device not found")
	PROCESS_NOT_FOUND        = errors.New("process not found using pid")
	UNABLET_TO_ATTACH_DEVICE = errors.New("unable to attach device")
	LOAD_SCRIPT_ERROR        = errors.New("unable to load script")
	CREATE_SCRIPT_ERROR      = errors.New("unable to create script")
)
