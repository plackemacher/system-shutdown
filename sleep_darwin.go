package system_shutdown

func Sleep() error {
	return invokeScript("tell application \"System Events\" to sleep")
}
