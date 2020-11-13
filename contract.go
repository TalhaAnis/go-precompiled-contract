package main

// #include <stdlib.h>
import "C"
import (
	"encoding/binary"
	"fmt"
	"sync"
	"time"
	"unsafe"
)

var mtx sync.Mutex

//export GetGasForData
func GetGasForData(ptr unsafe.Pointer, len C.int) uint64 {
	mtx.Lock()
	defer mtx.Unlock()
	return getGasForData(C.GoBytes(ptr, len))
}

//export Run
func Run(ptr unsafe.Pointer, len C.int) unsafe.Pointer {
	mtx.Lock()
	defer mtx.Unlock()
	rarr := run(C.GoBytes(ptr, len))
	cArr := C.CBytes(rarr)
	return cArr
}
func main() {}

/*///////////////////////////////////////////////////////////////////////////////
WARNING: DON'T MODIFY UPPER PART. QA TESTER WILL GENERATE AN ERROR AFTER SUBMSSTION
ONLY IMPORT SECTION CAN BE MODIFIED.
/////////////////////////////////////////////////////////////////////////////////*/

// getGasForData - Returns back gas required to execute the contract
func getGasForData([]byte) uint64 {
	// calculate gas here
	return uint64(5000000)
}

// run - Runs the contract, It recieve data as parsed byte and returns back a parsed byte array
func run(arr []byte) []byte {
	// Example of returning time in byte array
	msg := string(arr)
	fmt.Println(msg, "in go")
	timeBytes := make([]byte, 8)
	now := uint64(time.Now().UTC().UnixNano())
	binary.BigEndian.PutUint64(timeBytes, now)
	return timeBytes
}
