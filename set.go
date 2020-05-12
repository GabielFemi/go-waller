package go_waller


import (
	"fmt"
	"syscall"
	_ "syscall"
	"unsafe"
)

const (
	spiGetDesktopWallPaper = 0x0073
	spiSetDesktopWallPaper = 0x0014
	uiParam = 0x0000

	pifUpdateINIFile = 0x01
	spfSendChange    = 0x02

)
var (
	user32 = syscall.NewLazyDLL("user32.dll")
	systemParametersInfoW  = user32.NewProc("SystemParametersInfoW")
)
func SetWallPaper(filePath string) error{
	fileNameUTF16Ptr,  err := syscall.UTF16PtrFromString(filePath)
	if err != nil {
		return err
	}
	_, _, _ = systemParametersInfoW.Call(
		uintptr(spiSetDesktopWallPaper),
		uintptr(uiParam),
		uintptr(unsafe.Pointer(fileNameUTF16Ptr)),
		uintptr(pifUpdateINIFile|spfSendChange),
	)
	fmt.Println("Your wallpaper is now set.")
	return nil
}


