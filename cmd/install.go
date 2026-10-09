package cmd

import (
	"bufio"
	"os"
	"os/exec"
	"strings"
)

// currentBinary returns the path to the running executable, exiting after
// reporting the error if it can't be determined.
func currentBinary() string {
	bin, err := os.Executable()
	if err != nil {
		fatalf("%v", err)
	}
	return bin
}

// dispatchInstall asks whether to register the server with an agent and
// runs the matching installer, shared by init and the standalone commands.
func dispatchInstall(in *bufio.Scanner, env []envVar) {
	switch strings.ToLower(prompt(in, "Установить сервер для агента? [claude/codex/other/нет]", "нет")) {
	case "claude":
		installClaude(env)
	case "codex":
		installCodex(env)
	case "other", "другое", "другие":
		printSnippet(env)
	}
}

func runInstallCommand(name string, args []string) {
	c := exec.Command(name, args...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	if err := c.Run(); err != nil {
		fatalf("%s %s: %v", name, strings.Join(args, " "), err)
	}
}
