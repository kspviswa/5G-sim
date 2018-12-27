package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
)

var (
	helpCmd = &command{
		name:  "help",
		sname: "h",
		desp:  "displays help text",
		action: func(c *command, args []string) error {
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Cmd Name", "Cmd Shortname", "Flags", "Description", "Has Subcommands?", "Parent"})
			var cmdmap *map[string]*command
			cmdmap = &console.cmdmap
			parent := "NA"

			if len(args) > 0 && console.cmdmap[args[0]] != nil {
				cmdmap = &console.cmdmap[args[0]].subcommands
				parent = args[0]
			}

			for _, item := range *cmdmap {
				hasSubCmd := false
				if item.subcommands != nil {
					hasSubCmd = true
				}
				data := []string{
					item.name,
					item.sname,
					strings.Join(item.flags, ","),
					item.desp,
					strconv.FormatBool(hasSubCmd),
					parent,
				}

				table.Append(data)
			}

			table.Render()

			return nil
		},
	}

	shCmd = &command{
		name:  "shell",
		sname: "sh",
		desp:  "execute normal shell commands",
		action: func(c *command, args []string) error {
			cmdname := args[0]
			cmdargs := ""
			var out []byte
			var err error
			fmt.Println("length is %d", len(args))
			if len(args) > 1 {
				cmdargs = strings.Join(args[1:], " ")
				fmt.Print("args ")
				fmt.Println(cmdargs)
				out, err = exec.Command(cmdname, cmdargs).Output()
			} else {
				out, err = exec.Command(cmdname).Output()
			}

			if err != nil {
				fmt.Println("error occured")
				fmt.Println("%s", err.Error())
				return err
			}
			fmt.Println(string(out))
			return nil
		},
	}

	quitCmd = &command{
		name:  "quit",
		sname: "q",
		desp:  "Quit this simulator.",
		action: func(c *command, args []string) error {
			fmt.Println("Tearing down experiment... Please wait... !")
			fmt.Println("Wrapping up... just a min...")
			fmt.Println("Quit succeesfully")
			os.Exit(1)
			return nil
		},
	}

	loadCfgCmd = &command{
		name:  "loadcfg",
		sname: "loadcfg",
		desp:  "Load 5G Simulation Configuration from *.cfg file",
		action: func(c *command, args []string) error {
			path := args[0]
			err, nfs := loadFromFile(path)
			if err != nil {
				fmt.Println("Unable to load Config. Error : " + err.Error())
				return err
			} else {
				fmt.Println("Config loaded successfully")
			}
			nodes = *nfs
			return nil
		},
	}

	nfCmd = &command{
		name:        "nf",
		sname:       "nf",
		desp:        "Set of Network Function related command",
		subcommands: map[string]*command{"list": nfCmdList, "show": nfCmdShow},
		action: func(c *command, args []string) error {
			cmd, ok := c.subcommands[args[0]]
			if ok {
				cmd.action(cmd, args[1:])
				return nil
			}
			return nil //ToDO : return proper error
		},
	}

	nfCmdList = &command{
		name:  "list",
		sname: "ll",
		desp:  "List all registered network functions",
		action: func(c *command, args []string) error {
			nodeListAll()
			return nil
		},
	}

	nfCmdShow = &command{
		name:  "show",
		sname: "show",
		desp:  "Show detailed view of a specific network function",
		action: func(c *command, args []string) error {
			nodeShowDetails(args[0])
			return nil
		},
	}
)
