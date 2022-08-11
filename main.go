package main

import (
	"flag"
	"github.com/clarechu/quick-k8s/pkg/cmd"
	"github.com/spf13/pflag"
	"k8s.io/klog/v2"
	"os"
)

func init() {
	klog.InitFlags(nil)
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)

}

func main() {
	rootCmd := cmd.RootCommand(os.Args[1:])
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
