package repository

import (
	"context"
	"fmt"

	"net/smtp"

	help "github.com/adamnasrudin03/go-helpers"
	response_mapper "github.com/adamnasrudin03/go-helpers/response-mapper/v1"
	"github.com/adamnasrudin03/go-template/app/modules/message/dto"
)

func (r *messageRepo) SendEmail(ctx context.Context, params dto.SendEmailReq) (err error) {
	const opName = "MessageRepository-SendEmail"
	defer help.PanicRecover(opName)

	params.From = r.Cfg.Email.Sender
	err = params.Validate()
	if err != nil {
		r.Logger.Errorf("%v error validate: %v ", opName, err)
		return err
	}

	addr := fmt.Sprintf("%v:%v", r.Cfg.Email.Host, r.Cfg.Email.Port)
	auth := smtp.PlainAuth("", r.Cfg.Email.AuthEmail, r.Cfg.Email.AuthPassword, r.Cfg.Email.Host)
	err = smtp.SendMail(addr, auth, r.Cfg.Email.AuthEmail, append(params.To, params.Cc...), []byte(params.Body))
	if err != nil {
		r.Logger.Errorf("%v error: %v ", opName, err)
		return response_mapper.ErrFailedSendEmail()
	}

	return nil
}
