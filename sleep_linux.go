package system_shutdown

var sleepDbusArgs = []dbusArgs{
	{
		destination: "org.xfce.SessionManager",
		path:        "/org/xfce/SessionManager",
		iface:       "org.xfce.SessionManager",
		method:      "Suspend",
		body:        []interface{}{},
	},
	{
		destination: "org.freedesktop.login1",
		path:        "/org/freedesktop/login1",
		iface:       "org.freedesktop.login1.Manager",
		method:      "Suspend",
		body:        []interface{}{true},
	},
	{
		destination: "org.freedesktop.UPower",
		path:        "/org/freedesktop/UPower",
		iface:       "org.freedesktop.UPower",
		method:      "Suspend",
		body:        []interface{}{},
	},
	{
		destination: "org.freedesktop.Hal",
		path:        "/org/freedesktop/Hal/devices/computer",
		iface:       "org.freedesktop.Hal.Device.SystemPowerManagement",
		method:      "Suspend",
		body:        []interface{}{},
	},
}

func Sleep() (err error) {
	for _, v := range sleepDbusArgs {
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

	_, err = runCommand("systemctl", "suspend")
	return err
}
