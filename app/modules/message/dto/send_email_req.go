package dto

import (
	"strings"

	"github.com/adamnasrudin03/go-template/pkg/helpers"
)

type SendEmailReq struct {
	From    string   `json:"from"`
	To      []string `json:"to"`
	Cc      []string `json:"cc"`
	Subject string   `json:"subject"`
	Message string   `json:"message"`
	Body    string   `json:"body"`
}

func (m *SendEmailReq) Validate() error {
	err := m.ValidateTo()
	if err != nil {
		return err
	}
	if len(m.To) == 0 {
		return helpers.ErrIsRequired("Tujuan", "To")
	}

	err = m.ValidateCc()
	if err != nil {
		return err
	}

	m.Subject = strings.TrimSpace(m.Subject)
	m.Message = strings.TrimSpace(m.Message)
	if len(m.Message) == 0 {
		return helpers.ErrIsRequired("Pesan Surat", "Message Mail")
	}

	m.From = strings.TrimSpace(m.From)
	if len(m.From) == 0 {
		return helpers.ErrIsRequired("Pengirim", "From")
	}

	m.Body = "From: " + m.From + "\n" +
		"To: " + strings.Join(m.To, ",") + "\n" +
		"Cc: " + strings.Join(m.Cc, ",") + "\n" +
		"Subject: " + m.Subject + "\n\n" +
		m.Message
	return nil
}

func (m *SendEmailReq) ValidateTo() error {
	to := []string{}
	isToExist := map[string]bool{}
	for _, val := range m.To {
		val = strings.TrimSpace(val)
		if isToExist[val] {
			continue
		}

		if !helpers.IsValidEmail(val) {
			return helpers.ErrInvalidFormat("Tujuan Surel", "To Email")
		}

		isToExist[val] = true
		to = append(to, val)
	}
	m.To = to

	return nil
}

func (m *SendEmailReq) ValidateCc() error {

	cc := []string{}
	isCcExist := map[string]bool{}
	for _, val := range m.Cc {
		val = strings.TrimSpace(val)
		if isCcExist[val] {
			continue
		}

		if !helpers.IsValidEmail(val) {
			return helpers.ErrInvalidFormat("Cc Surel", "Cc Email")
		}

		isCcExist[val] = true
		cc = append(cc, val)
	}

	m.Cc = cc
	return nil
}
