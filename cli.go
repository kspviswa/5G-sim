package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type actionFunc func(c *command, args []string) error

type command struct {
	name        string
	sname       string
	desp        string
	flags       []string
	args        []string
	subcommands map[string]*command
	action      actionFunc
}

type cli struct {
	prompt string
	cmdmap map[string]*command
}

func (c *cli) doprompt() {
	fmt.Println()
	fmt.Print(c.prompt)
}

func (c *cli) addCommand(cmd *command) error {
	if cmd != nil {
		c.cmdmap[cmd.name] = cmd
	}
	return nil
}

func (c *cli) evaluatecmd(cmdstr string) (*command, bool) {
	cmd, ok := c.cmdmap[cmdstr]
	return cmd, ok
}

func (c *cli) processCommand(cmd *command) error {
	cmd.action(cmd, cmd.args)
	return nil
}

func (c *cli) initMemory() error {
	c.cmdmap = make(map[string]*command)
	return nil
}

func (c *cli) runner() {
	reader := bufio.NewScanner(os.Stdin)
	c.doprompt()
	for reader.Scan() {
		scan := reader.Text()
		tokens := strings.Split(scan, " ")
		cmd, ok := c.evaluatecmd(tokens[0])
		if ok {
			cmd.args = tokens[1:]
			c.processCommand(cmd)
		} else {
			if scan != "" {
				fmt.Println("Command not found. Try using `sh <cmd>` inorder to execute native commands")
			}
		}
		c.doprompt()
	}
}
