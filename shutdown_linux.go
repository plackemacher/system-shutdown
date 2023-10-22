package system_shutdown

var shutdownDbusArgs = []dbusArgs{
	{
		destination: "org.gnome.SessionManager",
		path:        "/org/gnome/SessionManager",
		iface:       "org.gnome.SessionManager",
		method:      "Shutdown",
		body:        []interface{}{},
	},
	{
		destination: "org.kde.ksmserver",
		path:        "/KSMServer",
		iface:       "org.kde.KSMServerInterface",
		method:      "logout",
		body:        []interface{}{-1, 2, 2},
	},
	{
		destination: "org.xfce.SessionManager",
		path:        "/org/xfce/SessionManager",
		iface:       "org.xfce.SessionManager",
		method:      "Shutdown",
		body:        []interface{}{true},
	},
	{
		destination: "org.freedesktop.login1",
		path:        "/org/freedesktop/login1",
		iface:       "org.freedesktop.login1.Manager",
		method:      "PowerOff",
		body:        []interface{}{true},
	},
	{
		destination: "org.freedesktop.PowerManagement",
		path:        "/org/freedesktop/PowerManagement",
		iface:       "org.freedesktop.PowerManagement",
		method:      "Shutdown",
		body:        []interface{}{},
	},
	{
		destination: "org.freedesktop.SessionManagement",
		path:        "/org/freedesktop/SessionManagement",
		iface:       "org.freedesktop.SessionManagement",
		method:      "Shutdown",
		body:        []interface{}{},
	},
	{
		destination: "org.freedesktop.ConsoleKit",
		path:        "/org/freedesktop/ConsoleKit/Manager",
		iface:       "org.freedesktop.ConsoleKit.Manager",
		method:      "Stop",
		body:        []interface{}{},
	},
	{
		destination: "org.freedesktop.Hal",
		path:        "/org/freedesktop/Hal/devices/computer",
		iface:       "org.freedesktop.Hal.Device.SystemPowerManagement",
		method:      "Shutdown",
		body:        []interface{}{},
	},
	{
		destination: "org.freedesktop.systemd1",
		path:        "/org/freedesktop/systemd1",
		iface:       "org.freedesktop.systemd1.Manager",
		method:      "PowerOff",
		body:        []interface{}{},
	},
}

func Shutdown() (err error) {
	for _, v := range shutdownDbusArgs {
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

	_, err = runCommand("shutdown", "-h", "now")
	return err
}
