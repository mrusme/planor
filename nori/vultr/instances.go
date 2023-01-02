package vultr

import (
	"context"
	// "errors"
	"fmt"
	"strings"

	"github.com/mrusme/planor/nori/models"

	"github.com/vultr/govultr/v2"
)

func (cloud *Vultr) ListInstances() ([]models.Instance, error) {
	input := govultr.ListOptions{}

	var instances []models.Instance
	ret, _, err := cloud.vultr.Instance.List(context.Background(), &input)
	if err != nil {
		return nil, err
	}

	for _, instance := range ret {
		newInstance := models.Instance{
			ID:   instance.ID,
			Name: instance.Label,

			Region: instance.Region,

			Type:         instance.Plan,
			Architecture: instance.Plan,
			CPUCores:     instance.VCPUCount,
			CPUThreads:   1,

			RAMSize:  instance.RAM,
			DiskSize: instance.Disk,

			OS:    instance.Os,
			Image: instance.ImageID,

			InternalIPv4: instance.InternalIP,

			IPv4:      instance.MainIP,
			NetmaskV4: instance.NetmaskV4,
			GatewayV4: instance.GatewayV4,

			IPv6:      instance.V6MainIP,
			NetworkV6: instance.V6Network,
			NetsizeV6: instance.V6NetworkSize,

			TransferLimit: instance.AllowedBandwidth,
			Info:          strings.Join(instance.Features, ", "),
			Status: fmt.Sprintf(
				"%s\nPower: %s\nServer: %s",
				instance.Status,
				instance.PowerStatus,
				instance.ServerStatus,
			),
			VNC: instance.KVM,
		}

		instances = append(instances, newInstance)
	}

	return instances, nil
}
