package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"sort"
	"strings"

	"github.com/charmbracelet/huh"
	"gopkg.in/yaml.v2"
)

type Cmd struct {
	Command []string
	Args    []string
}

type Filters []string
func (f *Filters) String() string {
	// Convert the slice to a string for flag.Value interface
	return fmt.Sprintf("%v", *f)
}

func (f *Filters) Set(value string) error {
	// Append the value to the slice
	*f = append(*f, value)
	return nil
}

func main() {

	//
	// Read flags
	//
	var filters Filters
	
	flag.Var(&filters, "f", "Command filter, must contain")
	flag.Parse()


	//
	// Read Config
	//

	// First check if config file is in current dir
	configFile := "config.yaml"
	_, err := os.Stat(configFile)
	if os.IsNotExist(err) {
		// check if config file is in well know location under the home folder
		usr, err := user.Current()
		must(err)

		configFile = usr.HomeDir + "/.config/cmdlist/config.yaml"

		_, err = os.Stat(configFile)
		if os.IsNotExist(err) {
			fmt.Printf("Config file not found, it is expected to exist here:\n%s\n", configFile)
			os.Exit(1)
		}
	}

	data, err := os.ReadFile(configFile)
	must(err)

	// Unmarshal the YAML into a map
	var commandMap map[string]Cmd
	err = yaml.Unmarshal(data, &commandMap)
	must(err)

	data, err = yaml.Marshal(commandMap)
	must(err)

	// Make a slice of strings from the map keys
	commandKeys := make([]string, 0, len(commandMap)) // Create a slice with initial capacity

	
	for key := range commandMap {
		// assume we will include the command
		include := true 

		if len(filters) > 0 {
			
			// it must match all the filters otherwise it is excluded
			for _,filter := range filters {
				if !strings.Contains(key, filter){
					include = false
				}
			}

		}

		if include {
			commandKeys = append(commandKeys, key)
		}
	}


	sort.Strings(commandKeys)

	var commandKey string

	form := huh.NewForm(

		huh.NewGroup(
			huh.NewSelect[string]().
				// Options(huh.NewOptions("nodes", "pods", "ls")...).
				Options(huh.NewOptions(commandKeys...)...).
				Title("Using config:  " + configFile).
				Description("\nChoose a command:").
				Height(20).
				Value(&commandKey),
		),
		// huh.NewGroup(
		// 	huh.NewSelect[string]().
		// 		Options(huh.NewOptions("start", "stop", "ssh")...).
		// 		Title("Choose:").
		// 		Value(&command),
		// ),
	)
	err = form.Run()
	must(err)

	{
		a := append(commandMap[commandKey].Command, commandMap[commandKey].Args...)
		fmt.Printf("command: %v\n", a)
		cmd := exec.Command(a[0], a[1:]...)

		// Redirect command's stdout/in/error to the current process's
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin

		//Inherit the current process's environment
		cmd.Env = os.Environ()

		err = cmd.Run()

	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
