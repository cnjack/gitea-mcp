package cmd

import (
	"context"
	"flag"
	"os"

	"gitea.com/gitea/gitea-mcp/operation"
	flagPkg "gitea.com/gitea/gitea-mcp/pkg/flag"
	"gitea.com/gitea/gitea-mcp/pkg/log"
)

var (
	transport string
	host      string
	port      int
	token     string

	debug *bool
)

func init() {
	flag.StringVar(
		&transport,
		"t",
		"stdio",
		"Transport type (stdio or sse)",
	)
	flag.StringVar(
		&transport,
		"transport",
		"stdio",
		"Transport type (stdio or sse)",
	)
	flag.StringVar(
		&host,
		"host",
		os.Getenv("GITEA_HOST"),
		"Gitea host",
	)
	flag.IntVar(
		&port,
		"port",
		8080,
		"sse port",
	)
	flag.StringVar(
		&token,
		"token",
		"",
		"Your personal access token",
	)
	flag.BoolFunc(
		"d",
		"debug mode (If -d flag is provided, debug mode will be enabled by default)",
		func(string) error {
			debug = new(bool)
			*debug = true
			return nil
		},
	)
	flag.BoolVar(
		&flagPkg.Insecure,
		"insecure",
		false,
		"ignore TLS certificate errors",
	)

	flag.Parse()

	flagPkg.Host = host
	if flagPkg.Host == "" {
		flagPkg.Host = "https://gitea.com"
	}

	flagPkg.Port = port

	flagPkg.Token = token
	if flagPkg.Token == "" {
		flagPkg.Token = os.Getenv("GITEA_ACCESS_TOKEN")
	}

	flagPkg.Mode = transport

	if debug != nil && *debug {
		flagPkg.Debug = *debug
	}
	if debug != nil && !*debug {
		flagPkg.Debug = os.Getenv("GITEA_DEBUG") == "true"
	}

	// Set insecure mode based on environment variable
	if os.Getenv("GITEA_INSECURE") == "true" {
		flagPkg.Insecure = true
	}
}

func Execute(version string) {
	defer log.Default().Sync()
	if err := operation.Run(transport, version); err != nil {
		if err == context.Canceled {
			log.Info("Server shutdown due to context cancellation")
			return
		}
		log.Fatalf("Run Gitea MCP Server Error: %v", err)
	}
}
