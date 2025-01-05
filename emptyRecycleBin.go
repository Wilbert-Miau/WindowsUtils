package main

import ("fmt"
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

func main(){
	EmptyRecycleBin()
}

func EmptyRecycleBin(){

	var handleToAWindow int =0
	var pszRootPath uintptr =uintptr(unsafe.Pointer(nil))
	flags:= SHERB_NOCONFIRMATION|SHERB_NOPROGRESSUI|SHERB_NOSOUND

	ret, _, err := procSHEmptyRecycleBinW.Call(uintptr(handleToAWindow),pszRootPath,uintptr(flags))
	if ret!=0{
		fmt.Println(err)
	}else{
		fmt.Println("recycle bin was emptied")
	}
}

// func EmptyRecycleBin(drivePath *string){
// 	var drivePathPointer *uint16
// 	if drivePath!=nil{
// 		result, err := syscall.UTF16PtrFromString(*drivePath)

// 		if err!=nil{
// 			return
// 		}
// 		drivePathPointer=result
// 	} else {
// 		drivePathPointer= nil

// 	}
// 	var handleToAWindow int =0
// 	// pinter
// 	var pszRootPath uintptr =uintptr(unsafe.Pointer(drivePathPointer))
// 	flags:= SHERB_NOCONFIRMATION|SHERB_NOPROGRESSUI|SHERB_NOSOUND

// 	ret, _, err := procSHEmptyRecycleBinW.Call(uintptr(handleToAWindow),pszRootPath,uintptr(flags))
// }
