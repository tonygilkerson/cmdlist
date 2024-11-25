package main

import (
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strings"

	"github.com/charmbracelet/huh"
	"gopkg.in/yaml.v2"
)

type Action struct {
	Command []string
	Args    []string
}

func main() {

	// Read the YAML file
	data, err := os.ReadFile("config.yaml")
	must(err)

	// Unmarshal the YAML into a map
	var actionMap map[string]Action
	err = yaml.Unmarshal(data, &actionMap)
	must(err)

	data, err = yaml.Marshal(actionMap)
	must(err)

	// Make a slice of strings from the map keys
	actionKeys := make([]string, 0, len(actionMap)) // Create a slice with initial capacity
	for key := range actionMap {
		actionKeys = append(actionKeys, key)
	}
  sort.Strings(actionKeys)
	
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
	must(err)

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

func must(err error) {
	if err != nil {
		panic(err)
	}
}
