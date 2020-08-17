// Package shell contains the `ww shell` command implementation.
package shell

import (
	"bytes"
	"context"
	"runtime"
	"text/template"

	"github.com/chzyer/readline"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"

	"github.com/spy16/sabre/repl"

	ww "github.com/wetware/ww/pkg"
	"github.com/wetware/ww/pkg/client"
	"github.com/wetware/ww/pkg/lang"
	anchorpath "github.com/wetware/ww/pkg/util/anchor/path"

	clientutil "github.com/wetware/ww/internal/util/client"
	ctxutil "github.com/wetware/ww/internal/util/ctx"
)

const bannerTemplate = `Wetware v{{.App.Version}}
Copyright {{.App.Copyright}}
Compiled with {{.GoVersion}} for {{.GOOS}}

`

var (
	root ww.Anchor = nopAnchor{} // see before()

	flags = []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "join",
			Aliases: []string{"j"},
			Usage:   "connect to cluster through specified peers",
			EnvVars: []string{"WW_JOIN"},
		},
		&cli.StringFlag{
			Name:    "discover",
			Aliases: []string{"d"},
			Usage:   "automatic peer discovery settings",
			Value:   "/mdns",
			EnvVars: []string{"WW_DISCOVER"},
		},
		&cli.StringFlag{
			Name:    "namespace",
			Aliases: []string{"ns"},
			Usage:   "cluster namespace (must match dial host)",
			Value:   "ww",
			EnvVars: []string{"WW_NAMESPACE"},
		},
		&cli.BoolFlag{
			Name:    "quiet",
			Aliases: []string{"q"},
			Usage:   "suppress banner message on interactive startup",
			EnvVars: []string{"WW_QUIET"},
		},
		&cli.BoolFlag{
			Name:    "dial",
			Usage:   "dial into a cluster using -join and -discover",
			EnvVars: []string{"WW_AUTODIAL"},
		},
	}
)

// Command constructor
func Command() *cli.Command {
	return &cli.Command{
		Name:   "shell",
		Usage:  "start an interactive REPL session",
		Flags:  flags,
		Before: before(),
		Action: run(),
	}
}

func run() cli.ActionFunc {
	return func(c *cli.Context) error {
		ww := lang.New(nil)

		if err := lang.BindAll(ww, root); err != nil {
			return err
		}

		lr, err := newLineReader(c)
		if err != nil {
			return err
		}
		defer lr.Close()

		return repl.New(ww,
			repl.WithBanner(banner(c)),
			repl.WithReaderFactory(readerFactory()),
			repl.WithPrompts("ww »", "   ›"),
			repl.WithInput(lr, nil),
			repl.WithOutput(c.App.Writer),
		).Loop(context.Background())
	}
}

// before the wetware client
func before() cli.BeforeFunc {
	return func(c *cli.Context) (err error) {
		if c.Bool("dial") {
			ctx := ctxutil.WithDefaultSignals(context.Background())
			root, err = clientutil.Dial(ctx, c)
		}

		return
	}
}

func after() cli.AfterFunc {
	return func(c *cli.Context) error {
		return root.(client.Client).Close()
	}
}

func readerFactory() repl.ReaderFactory {
	return repl.ReaderFactoryFunc(lang.NewReader)
}

func newLineReader(c *cli.Context) (r linereader, err error) {
	r.r, err = readline.NewEx(&readline.Config{
		HistoryFile: "/tmp/ww.tmp",
		Stdout:      c.App.Writer,
		Stderr:      c.App.ErrWriter,

		InterruptPrompt: "⏎",
		EOFPrompt:       "exit",

		/*
			TODO(enhancemenbt):  pass in the lang.Ww and configure autocomplete.
								 The lang.Ww instance will need to sup
		*/
		// AutoComplete: completer(ww),
	})

	return
}

func banner(c *cli.Context) string {
	if c.Bool("quiet") {
		return ""
	}

	return mustBanner(c)
}

func mustBanner(c *cli.Context) string {
	var buf bytes.Buffer

	templ := template.Must(template.New("banner").Parse(bannerTemplate))
	if err := templ.Execute(&buf, struct {
		*cli.Context
		GoVersion, GOOS string
	}{
		Context:   c,
		GoVersion: runtime.Version(),
		GOOS:      runtime.GOOS,
	}); err != nil {
		panic(err)
	}

	return buf.String()
}

type linereader struct {
	r *readline.Instance
}

func (l linereader) SetPrompt(s string) {
	l.r.SetPrompt(s)
}

func (l linereader) Readline() (line string, err error) {
	for {
		switch line, err = l.r.Readline(); err {
		case readline.ErrInterrupt:
			if len(line) == 0 {
				/* TODO(enhancement)

				- swallow ^C
				- clear the line & reset the prompt
				*/

				l.r.Clean()
				return "", nil
			}

			continue
		default:
			return // io.EOF
		}
	}
}

func (l linereader) Close() error {
	return l.r.Close()
}

type nopAnchor []string

func (a nopAnchor) String() string { return anchorpath.Join(a) }

func (a nopAnchor) Path() []string { return a }

func (nopAnchor) Ls(context.Context) ([]ww.Anchor, error) {
	// simulate missing host
	return nil, errors.Wrap(errors.New("open stream"),
		"failed to find any peer in table")
}

func (a nopAnchor) Walk(_ context.Context, path []string) ww.Anchor {
	return append(a, path...)
}
