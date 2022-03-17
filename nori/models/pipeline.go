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

