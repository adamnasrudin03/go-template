package helpers

import (
	"bytes"
	"context"
	"crypto/tls"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	IsStatusSuccess = map[int]bool{
		http.StatusOK:      true,
		http.StatusCreated: true,
	}
)

type CustomConfigClientRequest struct {
	Timeout int
}

func (c *CustomConfigClientRequest) SetTimeout(defaultTimeout int) int {
	if c.Timeout > 0 {
		return c.Timeout
	}

	return defaultTimeout
}

// GetHTTPRequestJSON ...
func GetHTTPRequestJSON(ctx context.Context, method string, url string, body io.Reader, headers ...map[string]string) (res []byte, statusCode int, err error) {
	defer PanicRecover("helpers-GetHTTPRequestJSON")
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	// iterate optional data of headers
	for _, header := range headers {
		for key, value := range header {
			req.Header.Set(key, value)
		}
	}

	timeout, _ := strconv.Atoi(os.Getenv("DEFAULT_TIMEOUT"))
	client := &http.Client{Timeout: time.Duration(timeout) * time.Second}
	r, err := client.Do(req)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	resp := StreamToByte(r.Body)

	defer func() {
		r.Body.Close()
		tags := map[string]interface{}{
			"http.headers":    req.Header,
			"http.method":     req.Method,
			"http.url":        req.URL.String(),
			"response.status": r.Status,
			"response.body":   string(resp),
		}
		log.Println(tags)
	}()

	return resp, r.StatusCode, nil
}

// GetHTTPRequestSkipVerify ...
func GetHTTPRequestSkipVerify(ctx context.Context, method string, url string, body io.Reader, customTimeOut int, headers ...map[string]string) (res []byte, statusCode int, err error) {
	defer PanicRecover("helpers-GetHTTPRequestSkipVerify")
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	// iterate optional data of headers
	for _, header := range headers {
		for key, value := range header {
			req.Header.Set(key, value)
		}
	}

	timeout, _ := strconv.Atoi(os.Getenv("DEFAULT_TIMEOUT"))
	if customTimeOut > 0 {
		timeout = customTimeOut
	}
	client := &http.Client{Timeout: time.Duration(timeout) * time.Second}
	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	r, err := client.Do(req)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	resp := StreamToByte(r.Body)
	defer func() {
		r.Body.Close()
		tags := map[string]interface{}{
			"http.headers":    req.Header,
			"http.method":     req.Method,
			"http.url":        req.URL.String(),
			"response.status": r.Status,
			"response.body":   string(resp),
		}
		log.Println(tags)
	}()

	return resp, r.StatusCode, nil
}

// StreamToString func
func StreamToString(stream io.Reader) string {
	if stream == nil {
		return ""
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.String()
}

// StreamToByte ...
func StreamToByte(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.Bytes()
}

// encode value param url
func QueryEscape(s string) string {
	return url.QueryEscape(strings.TrimSpace(s))
}
