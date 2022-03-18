package models

import (
  "time"
)

type PipelineStageAction struct {
  Name            string
  Status          string
  Summary         string
  PercentComplete int32
  UpdatedAt       time.Time
}

type PipelineStage struct {
  ID              string
  Name            string
  Status          string
  Actions         []PipelineStageAction
}

type Pipeline struct {
  ID              string
  Name            string
  Version         string
  Stages          []PipelineStage
  Error           error
  CreatedAt       time.Time
  UpdatedAt       time.Time
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

