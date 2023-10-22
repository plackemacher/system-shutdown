package system_shutdown

var hibernateDbusArgs = []dbusArgs{
	{
		destination: "org.xfce.SessionManager",
		path:        "/org/xfce/SessionManager",
		iface:       "org.xfce.SessionManager",
		method:      "Hibernate",
		body:        []interface{}{},
	},
	{
		destination: "org.freedesktop.login1",
		path:        "/org/freedesktop/login1",
		iface:       "org.freedesktop.login1.Manager",
		method:      "Hibernate",
		body:        []interface{}{true},
	},
	{
		destination: "org.freedesktop.UPower",
		path:        "/org/freedesktop/UPower",
		iface:       "org.freedesktop.UPower",
		method:      "Hibernate",
		body:        []interface{}{},
	},
	{
		destination: "org.freedesktop.Hal",
		path:        "/org/freedesktop/Hal/devices/computer",
		iface:       "org.freedesktop.Hal.Device.SystemPowerManagement",
		method:      "Hibernate",
		body:        []interface{}{},
	},
}

func Hibernate() (err error) {
	for _, v := range hibernateDbusArgs {
		if reply, err := dbusSend(
			v.destination,
			v.path,
			v.iface,
			v.method,
			v.body,
		); reply && err == nil {
			return nil
		}
	}

	_, err = runCommand("systemctl", "hibernate")
	return err
}
