// Copyright 2017 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

// openpitrix app server
package main

import (
	"flag"
	"fmt"
	"os"

	"openpitrix.io/openpitrix/pkg/cmd/app"
	config "openpitrix.io/openpitrix/pkg/config/app"
	"openpitrix.io/openpitrix/pkg/version"
)

func init() {
	// avoid glog warning
	// and skip -h or -help flag when call flag.Parse()
	flag.CommandLine.Usage = func() {
		// fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		// flag.PrintDefaults()
	}
	flag.CommandLine.Init(os.Args[0], flag.ContinueOnError)
	flag.CommandLine.Parse([]string{})
}

func main() {
	if len(os.Args) == 2 {
		if os.Args[1] == "-v" {
			fmt.Printf("openpitrix-app %s\n", version.ShortVersion)
			os.Exit(0)
		}
		if os.Args[1] == "-version" {
			fmt.Printf("openpitrix-app %s, build date %s\n", version.GitSha1Version, version.BuildDate)
			os.Exit(0)
		}

		if os.Args[1] == "-h" || os.Args[1] == "-help" {
			fmt.Println("Usage: [env=value] openpitrix-app")
			fmt.Println("       openpitrix-app -v")
			fmt.Println("       openpitrix-app -h")
			fmt.Println()

			config.PrintEnvs()
			fmt.Println()

			fmt.Println("See https://openpitrix.io/")
			os.Exit(0)
		}
	}

	app.Main(config.MustLoadConfig())
}
