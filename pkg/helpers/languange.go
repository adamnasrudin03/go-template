package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"golang.org/x/text/language"
)

const (
	Auto = "auto"
)

var (
	LangID = language.Indonesian.String()
	LangEn = language.English.String()
)

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
		return LangID
	}
	return s
}

func Translate(source, sourceLang, targetLang string) (string, error) {
	var (
		translation []interface{}
		text        []string
	)

	encodedSource := QueryEscape(source)
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
