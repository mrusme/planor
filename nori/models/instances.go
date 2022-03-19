package models

import (
  // "fmt"
)

type Instance struct {
  ID              string
  Name            string

  Type            string
  Architecture    string
  CPUCores        int
  CPUThreads      int

  Image           string
  IPv4            string
  IPv6            string
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

