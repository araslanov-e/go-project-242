package main

import (
	"code"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory; supports -r (recursive), -H (human-readable), -a (include hidden)",
		Description: "This utility prints the size of the specified file or directory. If no file is specified, the size of the current directory is displayed.\n\n" +
			"By default, the size of the specified directory is calculated without recursion and does not include hidden files and directories.\n" +
			"The -r flag enables recursive size calculation, while the -a flag includes hidden files and directories in the size calculation.\n" +
			"The -H flag formats the output in a human-readable format (e.g., 1.5KB, 2MB).",
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
			path := cmd.StringArg("path")
			if path == "" {
				path = "."
			}

			res, err := code.GetPathSize(path, cmd.Bool("recursive"), cmd.Bool("human"), cmd.Bool("all"))
			if err != nil {
				return fmt.Errorf("get path size %q: %w", path, err)
			}

			fmt.Printf("%s\t%s\n", res, path)

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
