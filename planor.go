package main

import (
	"flag"
	// "io/ioutil"
	// "log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mrusme/planor/nori"
	"github.com/mrusme/planor/ui"
	"github.com/mrusme/planor/ui/uictx"
)

func main() {
	var cloudProvider = "aws"
	var cloudProfile = ""

	flag.StringVar(&cloudProvider, "c", "aws",
		"cloud provider: aws, vultr, heroku")
	flag.StringVar(&cloudProfile, "p", "",
		"aws profile name, vultr/heroku/render/fleek api key env variable name")
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
