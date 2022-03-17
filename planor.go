package main

import (
  "fmt"
  "flag"
  "os"

  "github.com/mrusme/planor/nori"
)

func main() {
  var cloudProvider = "aws"
  var cloudProfile = "dev.verifyplus"

  flag.StringVar(&cloudProvider, "c", "aws", "cloud provider")
  flag.StringVar(&cloudProfile, "p", "", "cloud profile")
  flag.Parse()

  if cloudProvider == "" || cloudProfile == "" {
    flag.Usage()
    os.Exit(1)
  }

  cloud, err := nori.New(&cloudProvider, &cloudProfile)
  if err != nil {
    panic(err)
  }

  bla, err := cloud.ListPipelines()
  fmt.Printf("%v\n\n%v", bla, err)
}

