package fleek

import (
  "github.com/mrusme/planor/nori/models"

  // gofleek "github.com/mrusme/go-fleek"
)

func (cloud *Fleek) ListInstances() ([]models.Instance, error) {
  var instances []models.Instance
  ret, err := cloud.fleek.GetSitesByTeamId(cloud.teamId)
  if err != nil {
    return nil, err
  }

  for _, instance := range ret {
    newInstance := models.Instance{
      ID: instance.Slug,
      Name: instance.Name,

      Type: instance.Platform,

      Status: instance.PublishedDeploy.Status,
    }

    instances = append(instances, newInstance)
  }

  return instances, nil
}

