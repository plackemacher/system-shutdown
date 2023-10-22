package system_shutdown

func Shutdown() error {
	return invokeScript("tell application \"System Events\" to shut down")
}
