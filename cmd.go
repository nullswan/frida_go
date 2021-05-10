package frida_go

func Attach(targetPID uint) (sess *Session, err error) {
	d, _ := GetLocalDevice()
	return d.Attach(targetPID)
}
