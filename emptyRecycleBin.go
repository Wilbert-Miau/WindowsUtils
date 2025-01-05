package main

import ("fmt"
"os"
"syscall"
"unsafe")


const (
	//flags that supress dialogs, ui and sound
	SHERB_NOCONFIRMATION = 0x00000001
	SHERB_NOPROGRESSUI   = 0x00000002
	SHERB_NOSOUND        = 0x00000004
)

var (
	//dynamic loaded shell32.dll
	shell32 = syscall.NewLazyDLL("shell32.dll")
	
	procSHEmptyRecycleBinW = shell32.NewProc("SHEmptyRecycleBinW")
)

func EmptyRecycleBin(drivePath *string){
	var drivePathPointer *uint16
	if drivePath!=nil{
		result, err := syscall.UTF16PtrFromString(*drivePath)

		if err!=nil{
			return
		}
		drivePathPointer=result
	} else {
		drivePathPointer= nil

	}
	var handleToAWindow int =0
	// pinter
	var pszRootPath uintptr =uintptr(unsafe.Pointer(drivePathPointer))
	flags:= SHERB_NOCONFIRMATION|SHERB_NOPROGRESSUI|SHERB_NOSOUND

	ret, _, err := procSHEmptyRecycleBinW.Call(uintptr(handleToAWindow),pszRootPath,uintptr(flags))
}
