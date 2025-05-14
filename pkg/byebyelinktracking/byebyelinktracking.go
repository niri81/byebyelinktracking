package byebyelinktracking

import (
	"context"
	"log/slog"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/niri81/byebyelinktracking/internal"
	"golang.design/x/clipboard"
)

func Run(cfgFilePath string) {
	if strings.Contains(cfgFilePath, "$username") {
		user := os.Getenv("USER")
		if user == "" {
			slog.Error(
				"could not load your username, please use absolute path with your username or set $USER environment variable",
			)
		}

		cfgFilePath = strings.ReplaceAll(cfgFilePath, "$username", user)
	}

	ext := path.Ext(cfgFilePath)

	data, err := os.ReadFile(cfgFilePath)
	if err != nil {
		slog.Error("could not load config file", "err", err.Error())
		os.Exit(1)
	}

	config, err := parseConfig(data, ext)
	if err != nil {
		slog.Error("could not parse config", "err", err.Error())
		os.Exit(1)
	}

	slog.Info("successfully parsed config file", "filepath", cfgFilePath)

	err = clipboard.Init()
	if err != nil {
		slog.Error("could not access system clipboard", "err", err.Error())
		os.Exit(1)
	}

	ch := clipboard.Watch(context.Background(), clipboard.FmtText)
	var writeFromProgram bool

	for data := range ch {
		if writeFromProgram {
			writeFromProgram = false
			slog.Debug("skipping due to write from program")
			continue
		}

		strData := string(data)
		if _, err := url.ParseRequestURI(strData); err != nil {
			slog.Debug("skipping due to no URL data")
			continue
		}

		for _, entry := range config.Entries {
			if entry.MatchesHost(strData) {
				slog.Info("matched clipboard data on config file", "hosts", internal.ToString(entry.Host))

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
