package persistence

import (
	"fmt"
	"sync"

	"github.com/X4mp/coding-challenges/signing-service-challenge/domain"
	"github.com/google/uuid"
)

var ErrUnknownDevice error = fmt.Errorf("no device for ID registered")

// TODO: in-memory persistence ...
type Database struct {
	deviceMutexMapMutex       sync.Mutex
	deviceMutexMap            map[uuid.UUID]*sync.Mutex
	registeredDevicesMapMutex sync.RWMutex
	registeredDevicesMap      map[uuid.UUID]*domain.SignatureDevice
}

func NewDatabase() *Database {
	return &Database{
		registeredDevicesMap: make(map[uuid.UUID]*domain.SignatureDevice),
		deviceMutexMap:       make(map[uuid.UUID]*sync.Mutex),
	}
}

func (d *Database) LockDevice(deviceID uuid.UUID) func() {
	d.deviceMutexMapMutex.Lock()
	defer d.deviceMutexMapMutex.Unlock()

	deviceMutex, ok := d.deviceMutexMap[deviceID]
	if !ok {
		deviceMutex = &sync.Mutex{}
		d.deviceMutexMap[deviceID] = deviceMutex
	}
	deviceMutex.Lock()
	return func() {
		deviceMutex.Unlock()
	}
}

func (d *Database) IncrementCounter(deviceID uuid.UUID) {
	device, ok := d.registeredDevicesMap[deviceID]
	if ok {
		device.Counter += 1
		d.registeredDevicesMap[deviceID] = device
	}
}

func (d *Database) GetSignatureDevice(deviceId uuid.UUID) (device *domain.SignatureDevice, err error) {
	d.registeredDevicesMapMutex.RLock()
	defer d.registeredDevicesMapMutex.RUnlock()

	device, ok := d.registeredDevicesMap[deviceId]
	if !ok {
		err = ErrUnknownDevice
	}

	return
}

func (d *Database) StoreSignatureDevice(device *domain.SignatureDevice) {
	d.registeredDevicesMapMutex.Lock()
	defer d.registeredDevicesMapMutex.Unlock()

	d.registeredDevicesMap[device.DeviceId] = device
}
