package main

import (
	"fmt"
	"os"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/spf13/cobra"

	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	Built   string

	id, _ = os.Hostname()
)

var (
	_rootCmd = &cobra.Command{
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
		Use: "leo",
		Run: serve,
	}
)

func init() {
	_rootCmd.AddCommand(
		versionCommand(),
	)
}

func main() {
	_ = _rootCmd.Execute()
}

func versionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "show version info",
		Run: func(*cobra.Command, []string) {
			fmt.Printf(`LEO:
  Leo Server
    Version:   %s
    Built:     %s
`, Version, Built)
		},
	}
}

func newApp(gs *grpc.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Server(
			gs,
		),
	)
}

func serve(*cobra.Command, []string) {
	app, cleanup, err := wireApp()
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
