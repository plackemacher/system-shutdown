package system_shutdown

func invokeScript(script string) (err error) {
	_, err = runCommand("osascript", "-e", script)
	return err
}

func RequestPermissionDialog() error {
	return invokeScript("tell application \"System Events\" to stop current screen saver")
}
