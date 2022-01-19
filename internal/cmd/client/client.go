package client

import (
	"github.com/lthibault/log"
	"github.com/urfave/cli/v2"
	bootutil "github.com/wetware/ww/internal/util/boot"
	logutil "github.com/wetware/ww/internal/util/log"
	"github.com/wetware/ww/pkg/client"
)

var (
	logger log.Logger
	node   client.Node
)

func Command() *cli.Command {
	return &cli.Command{
		Name:  "client",
		Usage: "cli client for wetware clusters",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "discover",
				Aliases: []string{"d"},
				Usage:   "bootstrap discovery addr in URL-CIDR format",
				Value:   "tcp://127.0.0.1:8822/24", // TODO:  this should default to mudp
			},
			&cli.StringFlag{
				Name:    "ns",
				Usage:   "cluster namespace",
				Value:   "ww",
				EnvVars: []string{"WW_NS"},
			},
		},
		Subcommands: commands,
	}
}

var commands = []*cli.Command{
	Discover(),
	// Publish(),
	Subscribe(),
}

// ww client discover
func Discover() *cli.Command {
	return &cli.Command{
		Name:  "discover",
		Usage: "bootstrap client",
		Subcommands: []*cli.Command{
			Crawl(),
			Publish(),
		},
	}
}

// ww client subscribe <topic>
func Subscribe() *cli.Command {
	return &cli.Command{
		Name:    "subscribe",
		Aliases: []string{"sub"},
		Usage:   "subscribe to a pubsub topic",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "topic",
				Usage: "pubsub topic (\"\" is the cluster topic)",
			},
		},
		Before: before,
		Action: subscribe,
	}
}

func before(c *cli.Context) error {
	logger = logutil.New(c)

	crawler, err := bootutil.NewCrawler(c, logger)
	if err != nil {
		return err
	}

	node, err = client.DialDiscover(c.Context, crawler,
		client.WithLogger(logger),
		client.WithNamespace(c.String("ns")))
	return err
}
