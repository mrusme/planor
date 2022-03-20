package heroku

import (
  "context"
  // "errors"

  "github.com/mrusme/planor/nori/models"

  herokugo "github.com/heroku/heroku-go/v5"
)

func (cloud *Heroku) ListPipelines() ([]models.Pipeline, error) {
  apps, err := cloud.ListApps()
  if err != nil {
    return nil, err
  }

  var pipelines []models.Pipeline
  input := herokugo.ListRange{
    Field: "id",
    Max: 1,
    Descending: true,
  }
  for _, app := range apps {
    ret, err := cloud.heroku.BuildList(context.Background(), app.ID, &input)
    if err != nil {
      return pipelines, err
    }

    for _, build := range ret {
      newPipelineStage := models.PipelineStage {
        ID: build.ID,
        Name: build.Release.ID,
        Status: build.Status,
      }

      var stages []models.PipelineStage
      stages = append(stages, newPipelineStage)
      newPipeline := models.Pipeline{
        ID: app.ID,
        Name: app.Name,
        Stages: stages,
      }

      pipelines = append(pipelines, newPipeline)
    }
  }

  return pipelines, nil
}
