package models

import (
  // "fmt"
)

type Instance struct {
  ID              string
  Name            string
}

func (instance Instance) FilterValue() (string) {
  return instance.Name
}

func (instance Instance) Title() (string) {
  return instance.Name
}

func (instance Instance) Description() (string) {
  return instance.ID
}

