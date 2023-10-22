package system_shutdown

func Hibernate() error {
	return setSuspendState(true)
}
