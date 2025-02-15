package utils

import (
	"errors"
	"github.com/zero-remainder/go-ranker/internal/constants"
	"net/url"
	"regexp"
	"strings"
)

func ValidateURL(input string) error {
	if input == "" {
		return errors.New("URL is required")
	}
	if len(input) > 255 {
		return errors.New("URL must be at most 255 characters long")
	}
	if !strings.HasPrefix(input, constants.HTTP) && !strings.HasPrefix(input, constants.HTTPS) {
		input = constants.HTTPS + input
	}
	parsedURL, err := url.ParseRequestURI(input)
	if err != nil {
		return errors.New("invalid URL format")
	}
	domainPattern := `^[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(domainPattern)
	if !re.MatchString(parsedURL.Hostname()) {
		return errors.New("invalid domain in URL")
	}
	return nil
}
