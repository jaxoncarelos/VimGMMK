package LightingController

import (
	"fmt"
	"log"

	"github.com/google/gousb"
)

func Run() {
	ctx := gousb.NewContext()
	defer ctx.Close()
	ctx.OpenDevices(func(desc *gousb.DeviceDesc) bool {
		fmt.Printf("%s\n", desc)
		fmt.Printf("  Protocol: %s\n", desc.String())
		return false
	})
	dev, err := ctx.OpenDeviceWithVIDPID(0x320F, 0x5045)
	if err != nil {
		log.Fatalf("OpenDeviceWithVIDPID: %v", err)
	}
	defer dev.Close()
	// PID stands for
	fmt.Println("Device opened", dev)
	intf, done, err := dev.DefaultInterface()
	if err != nil {
		log.Fatalf("DefaultInterface: %v", err)
	}
	defer done()

	ep, err := intf.OutEndpoint(7)
	if err != nil {
		log.Fatalf("InEndpoint: %v", err)
	}
	data := make([]byte, 64)

	for i := range data {
		data[i] = byte(i)
	}
	numBytes, err := ep.Write(data)
	if numBytes != len(data) {
		log.Fatalf("Write: %v", err)
	}
	fmt.Println("Wrote", numBytes, "bytes")
}
