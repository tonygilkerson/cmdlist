package main

import (
	// "errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"github.com/charmbracelet/lipgloss"
	// xstrings "github.com/charmbracelet/x/exp/strings"
)

type GcloudCommand struct {
	Group string
	Workstation  string
	Project string
	Cluster string
	Config string
	Region string
	Account string
}

func main() {

	gcommands := map[string]GcloudCommand{
		"auth list": {Group: "auth", Workstation: ""},
		"ws2": {Group: "workstations", Workstation: "cws2"},
		"ws3": {Group: "workstations", Workstation: "cws3"},
	}

	var itemName string
	var command string

	form := huh.NewForm(

		huh.NewGroup(
			huh.NewSelect[string]().
				Options(huh.NewOptions("auth list","ws2","ws3","ws4")...).
				Title("Choose an item:").
				Value(&itemName),
		),
		huh.NewGroup(
			huh.NewSelect[string]().
				Options(huh.NewOptions("start","stop","ssh")...).
				Title("Choose an action for " + itemName + ":").
				Value(&command),
		),

	)

	err := form.Run()

	if err != nil {
		fmt.Println("Uh oh:", err)
		os.Exit(1)
	}

	prepareCmd := func() {
		time.Sleep(1 * time.Second)
	}

	_ = spinner.New().Title("Preparing your command...").Action(prepareCmd).Run()

	// Print order summary.
	{
		gcloudcmdExpanded := fmt.Sprintf("gcloud %s %s",gcommands[itemName].Group,command)

		var sb strings.Builder
		keyword := func(s string) string {
			return lipgloss.NewStyle().Foreground(lipgloss.Color("212")).Render(s)
		}
		fmt.Fprintf(&sb,
			"%s\n\n %s",
			lipgloss.NewStyle().Bold(true).Render("GCLOUD COMMAND"),
			keyword(gcloudcmdExpanded),
		)

		name := itemName
		if name != "" {
			name = ", " + name
		}
		fmt.Fprintf(&sb, "\n\nOk for? %s", name)


		fmt.Println(
			lipgloss.NewStyle().
				Width(40).
				BorderStyle(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("63")).
				Padding(1, 2).
				Render(sb.String()),
		)



		cmd := exec.Command("/Users/tgilkerson/google-cloud-sdk/bin/gcloud/" + gcloudcmdExpanded) 

		output, err := cmd.Output()
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println(string(output))
		}
	}
}

func getGcloudCommandOptions(){

}