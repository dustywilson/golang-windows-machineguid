package winmachguid

import (
	"syscall"
	"unsafe"
)

func GetWindowsMachineGuid() (guid string, err error) {
	var h syscall.Handle
	err = syscall.RegOpenKeyEx(syscall.HKEY_LOCAL_MACHINE, syscall.StringToUTF16Ptr(`SOFTWARE\Microsoft\Cryptography`), 0, syscall.KEY_READ, &h)
	if err != nil {
		return
	}
	defer syscall.RegCloseKey(h)
	var typ uint32
	var buf [74]uint16 // len = len(`{GUID-BLAH-BLAH}`) * 2
	n := uint32(len(buf))
	err = syscall.RegQueryValueEx(h, syscall.StringToUTF16Ptr("MachineGuid"), nil, &typ, (*byte)(unsafe.Pointer(&buf[0])), &n)
	if err != nil {
		return
	}
	guid = syscall.UTF16ToString(buf[:])
	return
}
