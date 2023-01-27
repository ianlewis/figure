/*
Copyright 2023 Google LLC.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
	"sigs.k8s.io/release-utils/version"
)

func checkExit(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func rootCmd() *cobra.Command {
	var font string
	var color string
	var printVer bool

	c := &cobra.Command{
		Use:     "figure [TEXT]",
		Short:   "Prints beautiful ASCII art from text.",
		Example: "figure --font isometric1 Go-Figure",
		RunE: func(cmd *cobra.Command, args []string) error {
			if printVer {
				v := version.GetVersionInfo()
				v.Name = cmd.Root().Name()
				v.Description = cmd.Root().Short
				fmt.Println(v.String())
				return nil
			}
			if len(args) > 0 {
				message := strings.Join(args, " ")
				fig := figure.NewFigure(message, font, false)
				if color != "" {
					fig = figure.NewColorFigure(message, font, color, false)
				}
				fig.Print()
			}
			return nil
		},
	}
	// c.AddCommand(verCmd)
	c.Flags().StringVarP(&font, "font", "f", "", "Font to use. Supports any font supported by go-figure.")
	c.Flags().StringVarP(&color, "color", "c", "", "Color to use. Supports any color supported by go-figure.")
	c.Flags().BoolVarP(&printVer, "version", "v", false, "Print version info and exit.")

	return c
}

func main() {
	checkExit(rootCmd().Execute())
}
