package vultr

import (
  "context"
  // "errors"

  "github.com/mrusme/planor/nori/models"

  "github.com/vultr/govultr/v2"
)

func (cloud *Vultr) ListInstances() ([]models.Instance, error) {
  input := govultr.ListOptions{
  }

  var instances []models.Instance
  ret, _, err := cloud.vultr.Instance.List(context.Background(), &input)
  if err != nil {
    return nil, err
  }

  for _, instance := range ret {
    newInstance := models.Instance{
      ID: instance.ID,
      Name: instance.Label,

      Type: instance.Plan,
      Architecture: instance.Plan,
      CPUCores: instance.VCPUCount,
      CPUThreads: 0,

      Image: instance.ImageID,
      IPv4: instance.MainIP,
      IPv6: instance.V6MainIP,
    }

    instances = append(instances, newInstance)
  }

  return instances, nil
}

