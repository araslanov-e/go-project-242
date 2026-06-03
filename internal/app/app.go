package app

import (
	"code"
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/urfave/cli/v3"
)

func Run(ctx context.Context, args []string, output io.Writer) error {
	return newCommand(output).Run(ctx, args)
}

type ExitError struct {
	Code int
	Err  error
}

func (e *ExitError) Error() string {
	return e.Err.Error()
}

func (e *ExitError) Unwrap() error {
	return e.Err
}

func ExitCode(err error) int {
	var exitErr *ExitError
	if errors.As(err, &exitErr) {
		return exitErr.Code
	}

	return 1
}

func newCommand(output io.Writer) *cli.Command {
	return &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory; supports -r (recursive), -H (human-readable), -a (include hidden)",
		Description: "This utility prints the size of the specified file or directory. If no file is specified, the size of the current directory is displayed.\n\n" +
			"By default, the size of the specified directory is calculated without recursion and does not include hidden files and directories.\n" +
			"The -r flag enables recursive size calculation, while the -a flag includes hidden files and directories in the size calculation.\n" +
			"The -H flag formats the output in a human-readable format (e.g., 1.5KB, 2MB).\n\n" +
			"Symbolic links are ignored and are not included in the calculated size.",
		ArgsUsage: "[file or directory]",
		Arguments: []cli.Argument{
			&cli.StringArg{
				Name: "path",
			},
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "recursive",
				Aliases: []string{"r"},
				Usage:   "recursive size of directories",
			},
			&cli.BoolFlag{
				Name:    "human",
				Aliases: []string{"H"},
				Usage:   "human-readable sizes (auto-select unit)",
			},
			&cli.BoolFlag{
				Name:    "all",
				Aliases: []string{"a"},
				Usage:   "include hidden files and directories",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			return run(output, cmd)
		},
	}
}

func run(output io.Writer, cmd *cli.Command) error {
	if cmd.NArg() > 0 {
		return &ExitError{
			Code: 2,
			Err:  errors.New("too many arguments: expected 0 or 1"),
		}
	}

	path := cmd.StringArg("path")
	if path == "" {
		path = "."
	}

	res, err := code.GetPathSize(path, cmd.Bool("recursive"), cmd.Bool("human"), cmd.Bool("all"))
	if err != nil {
		return &ExitError{
			Code: 1,
			Err:  fmt.Errorf("get path size %q: %w", path, err),
		}
	}

	_, err = fmt.Fprintf(output, "%s\t%s\n", res, path)
	if err != nil {
		return fmt.Errorf("write output: %w", err)
	}

	return nil
}
