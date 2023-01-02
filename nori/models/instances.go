package models

// "fmt"

type Instance struct {
	ID   string
	Name string

	Region string

	Type         string
	Architecture string
	CPUCores     int
	CPUThreads   int
	Hypervisor   string

	RAMSize  int // MB
	DiskSize int // GB

	OS    string
	Image string

	InternalIPv4 string

	IPv4      string
	NetmaskV4 string
	GatewayV4 string

	IPv6      string
	NetworkV6 string
	NetsizeV6 int

	TransferLimit int // GB

	Info   string
	Status string
	VNC    string
}

func (instance Instance) FilterValue() string {
	return instance.Name
}

func (instance Instance) Title() string {
	return instance.Name
}

func (instance Instance) Description() string {
	return instance.ID
}
