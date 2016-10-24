package main
/*
#include <stdlib.h>
extern void hego();
*/
import "C"
import (
	"fmt"
	"unsafe"
	//"reflect"
	"os"
)
func main(){
	fmt.Println("aiueo")
	//p := C.CString("point")
	//defer C.free(unsafe.Pointer(p))
	C.hego()
}
func assert(err error){
	if err != nil{
		fmt.Printf("---Error=%s", err.Error)
		os.Exit(1)
	}
}