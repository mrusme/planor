package main

import (
  "flag"
  // "io/ioutil"
  // "log"
  "os"

  "github.com/mrusme/planor/nori"
  "github.com/mrusme/planor/ui"
  "github.com/mrusme/planor/ui/uictx"
  tea "github.com/charmbracelet/bubbletea"
)

func main() {
  var cloudProvider = "aws"
  var cloudProfile = "dev.verifyplus"

  flag.StringVar(&cloudProvider, "c", "aws", "cloud provider: aws, vultr")
  flag.StringVar(&cloudProfile, "p", "", "aws profile name, vultr api key env variable name")
  flag.Parse()

  if cloudProvider == "" || cloudProfile == "" {
    flag.Usage()
    os.Exit(1)
  }

  cloud, err := nori.New(&cloudProvider, &cloudProfile)
  if err != nil {
    panic(err)
  }

  ctx := uictx.New(&cloud)

  // log.SetOutput(ioutil.Discard)

  tui := tea.NewProgram(ui.NewModel(&ctx), tea.WithAltScreen())
  err = tui.Start()
  if err != nil {
    panic(err)
  }

}

