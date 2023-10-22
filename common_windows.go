package system_shutdown

import (
	"errors"
	"golang.org/x/sys/windows"
	"syscall"
)

//goland:noinspection GoSnakeCaseUsage
var SE_SHUTDOWN_NAME *uint16

func requestPrivileges() (err error) {
	var token windows.Token
	err = windows.OpenProcessToken(
		windows.CurrentProcess(),
		windows.TOKEN_ADJUST_PRIVILEGES|windows.TOKEN_QUERY,
		&token,
	)
	if err != nil {
		return err
	}

	var tkp windows.Tokenprivileges
	err = windows.LookupPrivilegeValue(nil, SE_SHUTDOWN_NAME, &tkp.Privileges[0].Luid)
	if err != nil {
		return err
	}

	tkp.PrivilegeCount = 1
	tkp.Privileges[0].Attributes = windows.SE_PRIVILEGE_ENABLED
	err = windows.AdjustTokenPrivileges(token, false, &tkp, 0, nil, nil)
	return err
}

func exitWindows(flag uint32) (err error) {
	err = requestPrivileges()
	if err != nil {
		return err
	}

	err = windows.ExitWindowsEx(
		flag|windows.EWX_FORCEIFHUNG,
		windows.SHTDN_REASON_MAJOR_OPERATINGSYSTEM|windows.SHTDN_REASON_MINOR_UPGRADE|windows.SHTDN_REASON_FLAG_PLANNED,
	)
	return err
}

func setSuspendState(hibernate bool) (err error) {
	err = requestPrivileges()
	if err != nil {
		return err
	}

	powerProf := syscall.MustLoadDLL("PowrProf.dll")
	defer func(powerProf *syscall.DLL) {
		localErr := powerProf.Release()
		err = errors.Join(err, localErr)
	}(powerProf)

	SetSuspendState := powerProf.MustFindProc("SetSuspendState")
	var hibernateVal uintptr
	if hibernate {
		hibernateVal = 1
	} else {
		hibernateVal = 0
	}

	var result uintptr
	result, _, err = SetSuspendState.Call(hibernateVal, 0, 0)
	if result == 0 {
		return err
	}
	return nil
}

func init() {
	var err error
	SE_SHUTDOWN_NAME, err = windows.UTF16PtrFromString("SeShutdownPrivilege")
	if err != nil {
		panic(err)
	}
}
