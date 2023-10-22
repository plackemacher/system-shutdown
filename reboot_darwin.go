package system_shutdown

func Reboot() error {
	return invokeScript("tell application \"System Events\" to restart")
}
