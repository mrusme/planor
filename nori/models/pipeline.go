package models

import (
)

type PipelineStage struct {
  ID              string
  Name            string
  Status          string
}

type Pipeline struct {
  ID              string
  Name            string
  Version         string
  Stages          []PipelineStage
  Error           error
}

func (pipeline Pipeline) FilterValue() (string) {
  return pipeline.Name
}

func (pipeline Pipeline) Title() (string) {
  return pipeline.Name
}

func (pipeline Pipeline) Description() (string) {
  if len(pipeline.Stages) > 0 {
    return pipeline.Stages[0].Status
  }

  return "N/A"
}

