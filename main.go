package main

import (
	// "errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	// "time"

	"github.com/charmbracelet/huh"
	// "github.com/charmbracelet/huh/spinner"
	// "github.com/charmbracelet/lipgloss"
	// xstrings "github.com/charmbracelet/x/exp/strings"
)

type Action struct {
	Command []string
	Args    []string
}

func main() {

	actions := map[string]Action{
		"nodes": {
			Command: []string{"kubectl"},
			Args:    []string{"get", "nodes"},
		},
		"pods": {
			Command: []string{"kubectl"},
			Args:    []string{"get", "pods"},
		},
		"lstest": {
			Command: []string{"ls"},
			Args:    []string{"-l"},
		},
	}

	var actionKey string

	form := huh.NewForm(

		huh.NewGroup(
			huh.NewSelect[string]().
				Options(huh.NewOptions("nodes", "pods", "ls")...).
				Title("Choose an item:").
				Value(&actionKey),
		),
		// huh.NewGroup(
		// 	huh.NewSelect[string]().
		// 		Options(huh.NewOptions("start", "stop", "ssh")...).
		// 		Title("Choose an action for "+actionKey+":").
		// 		Value(&command),
		// ),
	)

	err := form.Run()

	if err != nil {
		fmt.Println("Uh oh:", err)
		os.Exit(1)
	}

	// prepareCmd := func() {
	// 	time.Sleep(800 * time.Microsecond)
	// }

	// _ = spinner.New().Title("Preparing your command...").Action(prepareCmd).Run()

	{
		a := append(actions[actionKey].Command, actions[actionKey].Args...)
		fmt.Printf("command: %v\n", a)
		cmd := exec.Command(a[0], a[1:]...)

		stdout := new(strings.Builder)
		stderr := new(strings.Builder)
		cmd.Stdout = stdout
		cmd.Stderr = stderr
		
		//Inherit the current process's environment
		cmd.Env = os.Environ()
		
		err := cmd.Run()

		// output, err := cmd.Output()
		if err != nil {
			fmt.Println("Error:", err)
		} else {

			// fmt.Println(string(output))
			fmt.Println(stderr.String())
			fmt.Println(stdout.String())
		}

	}
}

func getCommandOptions() {

}
