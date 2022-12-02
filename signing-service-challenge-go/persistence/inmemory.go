package persistence

import (
	"fmt"
	"sync"

	"github.com/X4mp/coding-challenges/signing-service-challenge/domain"
	"github.com/google/uuid"
)

var UnknownDeviceError error = fmt.Errorf("No device for ID registered")

// TODO: in-memory persistence ...
type Database struct {
	deviceCounterMapMutex sync.Mutex
	deviceCounterMap map[uuid.UUID]uint64
	registeredDevicesMapMutex sync.RWMutex
	registeredDevicesMap map[uuid.UUID]*domain.SignatureDevice
}

func NewDatabase() *Database {
	return &Database{
		deviceCounterMap: make(map[uuid.UUID]uint64),
		registeredDevicesMap: make(map[uuid.UUID]*domain.SignatureDevice),
	}
}

func (d *Database) IncCounter(deviceID uuid.UUID) {
	d.deviceCounterMapMutex.Lock()
	defer d.deviceCounterMapMutex.Unlock()

	counter, ok := d.deviceCounterMap[deviceID]
	if ok {
		counter += 1
		d.deviceCounterMap[deviceID] = counter
	}
}

func (d *Database) GetSignatureDevice(deviceId uuid.UUID) (device *domain.SignatureDevice, err error) {
	d.registeredDevicesMapMutex.RLock()
	defer d.registeredDevicesMapMutex.Unlock()

	device, ok := d.registeredDevicesMap[deviceId]
	if !ok {
		err = UnknownDeviceError
	}

	return
}