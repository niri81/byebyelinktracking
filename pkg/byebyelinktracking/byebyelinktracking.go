package byebyelinktracking

import (
	"context"
	"fmt"
	"log/slog"
	"net/url"
	"os"

	"golang.design/x/clipboard"
)

func Run(cfgFile string) {
	data, err := os.ReadFile(cfgFile)
	if err != nil {
		slog.Error("could not load config file", "err", err.Error())
		os.Exit(1)
	}

	config := parseConfig(data)

	err = clipboard.Init()
	if err != nil {
		slog.Error("could not access system clipboard", "err", err.Error())
		os.Exit(1)
	}

	ch := clipboard.Watch(context.Background(), clipboard.FmtText)
	var writeFromProgram bool

	for data := range ch {
		if writeFromProgram {
			slog.Debug("skipping due to write from program")
			continue
		}

		strData := string(data)
		if _, err := url.Parse(strData); err != nil {
			slog.Debug("skipping due to no URL data")
			continue
		}

		for _, entry := range config.Entries {
			if entry.MatchesHost(strData) {
				slog.Info("matched clipboard data on config file", "hosts", fmt.Sprint(entry.Host))

				newData, ok := entry.Handle(strData)
				if ok && strData != newData {
					writeFromProgram = true
					strData = newData
				}
			}
		}

		slog.Debug("Writing to clipboard", "data", strData)
		clipboard.Write(clipboard.FmtText, []byte(strData))
	}
}
