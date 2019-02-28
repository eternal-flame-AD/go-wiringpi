package i2c

// #cgo LDFLAGS: -lwiringPi
// #include <string.h>
// #include <wiringPiI2C.h>
import "C"
import (
	"errors"
	"sync"
)

var retCodeToErrLock = new(sync.Mutex)

func retCodeToErr(ret int) error {
	if ret >= 0 {
		return nil
	}
	retCodeToErrLock.Lock()
	defer retCodeToErrLock.Unlock()
	return errors.New(C.GoString(C.strerror(C.int(ret))))
}

// Open a device of given ID and returns the file descriptor
func Open(devID uint8) (fd int, err error) {
	res := int(C.wiringPiI2CSetup(C.int(devID)))
	if res < 0 {
		return 0, retCodeToErr(res)
	}
	return res, nil
}

// Read is simple device read
func Read(fd int) (uint8, error) {
	res := int(C.wiringPiI2CRead(C.int(fd)))
	if res < 0 {
		return 0, retCodeToErr(res)
	}
	return uint8(res), nil
}

// Write is simple device write
func Write(fd int, data uint8) error {
	res := int(C.wiringPiI2CWrite(C.int(fd), C.int(data)))
	return retCodeToErr(res)
}

// WriteReg8 writes 8-bit data to device register
func WriteReg8(fd int, reg int, data uint8) error {
	res := int(C.wiringPiI2CWriteReg8(C.int(fd), C.int(reg), C.int(data)))
	return retCodeToErr(res)
}

// WriteReg16 writes 16-bit data to device register
func WriteReg16(fd int, reg int, data uint16) error {
	res := int(C.wiringPiI2CWriteReg16(C.int(fd), C.int(reg), C.int(data)))
	return retCodeToErr(res)
}

// ReadReg8 reads 8-bit data from device register
func ReadReg8(fd int, reg int) (uint8, error) {
	res := int(C.wiringPiI2CReadReg8(C.int(fd), C.int(reg)))
	if res < 0 {
		return 0, retCodeToErr(res)
	}
	return uint8(res), nil
}

// ReadReg16 reads 16-bit data from device register
func ReadReg16(fd int, reg int) (uint16, error) {
	res := int(C.wiringPiI2CReadReg16(C.int(fd), C.int(reg)))
	if res < 0 {
		return 0, retCodeToErr(res)
	}
	return uint16(res), nil
}
