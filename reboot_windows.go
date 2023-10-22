package system_shutdown

import "golang.org/x/sys/windows"

func Reboot() error {
	return exitWindows(windows.EWX_REBOOT)
}
