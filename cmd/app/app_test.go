package app

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/urfave/cli/v3"
)

func TestCommandExposesExpectedHierarchy(t *testing.T) {
	t.Parallel()

	root := Command()
	if root.Name != "app" {
		t.Fatalf("root command name = %q, want %q", root.Name, "app")
	}

	start := requireCommand(t, root, "start")
	if start.Action == nil {
		t.Fatal("start command has no action")
	}

	database := requireCommand(t, root, "db")
	for _, commandName := range []string{
		"init",
		"migrate",
		"rollback",
		"lock",
		"unlock",
		"create_go",
		"create_sql",
		"status",
		"mark_applied",
	} {
		command := requireCommand(t, database, commandName)
		if command.Action == nil {
			t.Fatalf("%s command has no action", commandName)
		}
	}
}

func TestCommandRendersHelp(t *testing.T) {
	t.Parallel()

	var output bytes.Buffer
	root := Command()
	root.Writer = &output
	root.ErrWriter = &output

	if err := root.Run(context.Background(), []string{"app", "--help"}); err != nil {
		t.Fatalf("render help: %v", err)
	}

	help := output.String()
	for _, expected := range []string{"COMMANDS:", "start", "db"} {
		if !strings.Contains(help, expected) {
			t.Errorf("help output does not contain %q:\n%s", expected, help)
		}
	}
}

func requireCommand(t *testing.T, parent *cli.Command, name string) *cli.Command {
	t.Helper()

	command := parent.Command(name)
	if command == nil {
		t.Fatalf("%s command was not registered", name)
	}
	return command
}
