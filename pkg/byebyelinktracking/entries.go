package byebyelinktracking

import (
	"log/slog"
	"net/url"
	"regexp"
)

type configEntry struct {
	Host           []regexp.Regexp `json:"hosts"`
	Parameter_name []string        `json:"param_name"`
	Regex          []regexp.Regexp `json:"regex"`
}

type ConfigEntryInterface interface {
	MatchesHost(url string) bool
	Handle(url string) (string, bool)
}

func (cep *configEntry) MatchesHost(url string) bool {
	for _, exp := range cep.Host {
		if exp.MatchString(url) {
			return true
		}
	}

	return false
}

func (cep *configEntry) Handle(input_url string) (string, bool) {
	for _, exp := range cep.Regex {
		input_url = exp.ReplaceAllString(input_url, "")
	}

	url_parsed, err := url.Parse(input_url)
	if err != nil {
		slog.Error("could not parse URL", "err", err.Error())
		return "", false
	}

	query := url_parsed.Query()
	for _, i := range cep.Parameter_name {
		query.Del(i)
	}

	url_parsed.RawQuery = query.Encode()

	return url_parsed.String(), true
}
