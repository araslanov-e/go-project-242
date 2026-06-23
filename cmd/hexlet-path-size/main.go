package main

import (
	"code/internal/app"
	"context"
	"fmt"
	"os"
)

func main() {
	if err := app.Run(context.Background(), os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
}
