package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/text/language"
)

const (
	Auto = "auto"
)

// javascript "encodeURI()"
// so we embed js to our golang program
func EncodeURI(s string) string {
	return url.QueryEscape(s)
}

func defaultSourceLang(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return Auto
	}
	return s
}

func defaultTargetLang(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return language.Indonesian.String()
	}
	return s
}

func Translate(source, sourceLang, targetLang string) (string, error) {
	var (
		translation []interface{}
		text        []string
	)

	encodedSource := url.QueryEscape(source)
	url := fmt.Sprintf("https://translate.googleapis.com/translate_a/single?client=gtx&sl=%s&tl=%s&dt=t&q=%s",
		defaultSourceLang(sourceLang), defaultTargetLang(targetLang), encodedSource)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(body, &translation)
	if err != nil {
		return "", err
	}

	if len(translation) > 0 {
		inner := translation[0]
		for _, slice := range inner.([]interface{}) {
			for _, translatedText := range slice.([]interface{}) {
				text = append(text, fmt.Sprintf("%v", translatedText))
				break
			}
		}
		return strings.Join(text, ""), nil
	}

	return "", errors.New("no translated data in response")
}
