package system_shutdown

func Sleep() error {
	return setSuspendState(false)
}
