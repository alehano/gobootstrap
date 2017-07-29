/*
Commands that can run via command line by flag -run "name"

To register command you have to create new command with:
	* RegisterCLI("name", "description", func()error)
	* RegisterCLIRunner("command_name", "description", CLIRunner)

CLIRunner - is an interface with CLIRun()error

*/
package cli

import (
	"log"
	"sort"
	"errors"
	"fmt"
	"time"
	"flag"
	"os"
)

// Command Store
var store map[string]command = make(map[string]command)

// Interface for run function
type Runner interface {
	CLIRun() error
}

// Command
type command struct {
	name  string
	descr string
	fn    []Runner
}

// Get name and description
func (c command) getInfo() string {
	return fmt.Sprintf("%q - %s", c.name, c.descr)
}

// Wrapper Runner to register plain function
type wRunner struct {
	fn func() error
}

func (r wRunner) CLIRun() error  {
	return r.fn()
}

// Init new command
func RegisterCLIRunner(name, descr string, rn ...Runner) {
	cmd := command{name, descr, []Runner{}}
	if len(rn) > 0 {
		for _, r := range rn {
			cmd.fn = append(cmd.fn, r)
		}
	}
	store[name] = cmd
}

func RegisterCLI(name, descr string, fn func()error) {
	rn := wRunner{fn: fn}
	RegisterCLIRunner(name, descr, rn)
}

// Get sorted commands
func GetAllCommandsInfo() []string {
	res := []string{}
	for _, cmd := range store {
		res = append(res, cmd.getInfo())
	}
	sort.Strings(res)
	return res
}

// Runs command by a name
func Run(name string) {
	RunVerbose(name, func(info string) {}, func(err error) {
		log.Println(err)
	})
}

func RunVerbose(name string, logFn func(string), errFn func(error)) {
	if cmd, ok := store[name]; !ok {
		errFn(errors.New(fmt.Sprintf("Can't find CLI Command: %q. Check if package with CLI included to import in main.go", name)))
	} else {
		if len(cmd.fn) == 0 {
			errFn(errors.New(fmt.Sprintf("No functions were registered for CLI Command: %q", name)))
		}
		start := time.Now()
		logFn(fmt.Sprintf("Start CLI Command: %s", cmd.getInfo()))
		for _, f := range cmd.fn {
			err := f.CLIRun()
			if err != nil {
				errFn(errors.New(fmt.Sprintf("CLI Command: %q error: %s", cmd.name, err)))
			}
		}
		logFn(fmt.Sprintf("Finish CLI Command: %q, for: %s", cmd.name, time.Since(start)))
	}
}

func CheckAndRun() {
	commandInfo := "Run command by a name\n\tE.g. -run version\n\t-----------------\n"
	for _, info := range GetAllCommandsInfo() {
		commandInfo += fmt.Sprintf("\t%s\n", info)
	}
	var runFlag = flag.String("run", "", commandInfo)
	flag.Parse()
	if *runFlag != "" {
		Run(*runFlag)
		os.Exit(0)
	}
}