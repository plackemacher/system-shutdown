package system_shutdown

import "golang.org/x/sys/windows"

func Shutdown() error {
	return exitWindows(windows.EWX_SHUTDOWN)
}
