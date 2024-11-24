package main

import (
	// "errors"
	"fmt"
	"log"

	// "maps"
	"os"
	"os/exec"
	"strings"

	// "time"

	"github.com/charmbracelet/huh"
	"gopkg.in/yaml.v2"
	// "gopkg.in/yaml.v2"
	// "github.com/charmbracelet/huh/spinner"
	// "github.com/charmbracelet/lipgloss"
	// xstrings "github.com/charmbracelet/x/exp/strings"
)

type Action struct {
	Command []string
	Args    []string
}

func main() {

	// Read the YAML file
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}


	// Unmarshal the YAML into a map
	var actionMap map[string]Action
	err = yaml.Unmarshal(data, &actionMap)
	if err != nil {
		panic(err)
	}

	// actionMap := map[string]Action{
	// 	"nodes": {
	// 		Command: []string{"kubectl"},
	// 		Args:    []string{"get", "nodes"},
	// 	},
	// 	"pods": {
	// 		Command: []string{"kubectl"},
	// 		Args:    []string{"get", "pods"},
	// 	},
	// 	"node \twide": {
	// 		Command: []string{"kubectl"},
	// 		Args:    []string{"get", "nodes", "-owide"},
	// 	},
	// 	"lstest": {
	// 		Command: []string{"ls"},
	// 		Args:    []string{"-l"},
	// 	},
	// }




	data, err = yaml.Marshal(actionMap)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("configtest.yaml", data, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Make a slice of strings from the map keys
	actionKeys := make([]string, 0, len(actionMap)) // Create a slice with initial capacity
	for key := range actionMap {
		actionKeys = append(actionKeys, key)
	}

	var actionKey string

	form := huh.NewForm(

		huh.NewGroup(
			huh.NewSelect[string]().
				// Options(huh.NewOptions("nodes", "pods", "ls")...).
				Options(huh.NewOptions(actionKeys...)...).
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

	err = form.Run()

	if err != nil {
		fmt.Println("Uh oh:", err)
		os.Exit(1)
	}

	// prepareCmd := func() {
	// 	time.Sleep(800 * time.Microsecond)
	// }

	// _ = spinner.New().Title("Preparing your command...").Action(prepareCmd).Run()

	{
		a := append(actionMap[actionKey].Command, actionMap[actionKey].Args...)
		fmt.Printf("command: %v\n", a)
		cmd := exec.Command(a[0], a[1:]...)

		stdout := new(strings.Builder)
		stderr := new(strings.Builder)
		cmd.Stdout = stdout
		cmd.Stderr = stderr

		//Inherit the current process's environment
		cmd.Env = os.Environ()

		err = cmd.Run()

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
