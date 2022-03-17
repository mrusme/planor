package main

import (
  "flag"
  "io/ioutil"
  "log"
  "os"

  "github.com/mrusme/planor/ui"
  tea "github.com/charmbracelet/bubbletea"
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

  // cloud, err := nori.New(&cloudProvider, &cloudProfile)
  // if err != nil {
  //   panic(err)
  // }

  log.SetOutput(ioutil.Discard)

  tui := tea.NewProgram(ui.NewModel(), tea.WithAltScreen())
  err := tui.Start()
  if err != nil {
    panic(err)
  }

  // bla, err := cloud.ListPipelines()
  // fmt.Printf("%v\n\n%v", bla, err)
}

