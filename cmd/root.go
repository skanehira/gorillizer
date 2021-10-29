package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var (
	zero = "ウ"
	one  = "ホ"
)

var rootCmd = &cobra.Command{
	Use: "gorillizer",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func exitError(msg interface{}) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}

func gorillize(args []string) string {
	var out []string
	for _, a := range args {
		for _, b := range []byte(a) {
			b := fmt.Sprintf("%b", b)
			b = strings.ReplaceAll(b, "0", zero)
			b = strings.ReplaceAll(b, "1", one)
			out = append(out, b)
		}
	}
	return strings.Join(out, " ")
}

func humanize(args []string) string {
	out := []byte{}
	for _, a := range args {
		a = strings.ReplaceAll(a, zero, "0")
		a = strings.ReplaceAll(a, one, "1")
		h, err := strconv.ParseInt(a, 2, 0) // convert to hex
		if err != nil {
			return err.Error()
		}
		out = append(out, byte(h))
	}
	return string(out)
}

func Execute() {
	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			_ = cmd.Help()
		}

		if args[0] == "humanize" {
			if len(args) == 1 {
				_ = cmd.Help()
				return
			}
			args = args[1:]
			fmt.Println(humanize(args))
			return
		}
		fmt.Println(gorillize(args))
	}

	if err := rootCmd.Execute(); err != nil {
		exitError(err)
	}
}
