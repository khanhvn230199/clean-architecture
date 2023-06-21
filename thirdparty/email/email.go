package email

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net/smtp"

	errorpkg "github.com/example-golang-projects/clean-architecture/packages/error"
)

type SMTPConfig struct {
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
	Encrypt     string `yaml:"encrypt"`
	FromAddress string `yaml:"from_address"`
}

func (c *SMTPConfig) SMTPServer() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

type Client struct {
	cfg *SMTPConfig
}

// New ...
func New(cfg *SMTPConfig) *Client {
	//cfg.Encrypt = strings.ToLower(cfg.Encrypt)
	c := &Client{
		cfg: cfg,
	}
	return c
}

type SendEmailCommand struct {
	FromName    string
	ToAddresses []string
	Subject     string
	Content     string
}

func (c *Client) Ping() error {
	client, err := c.Dial()
	if err != nil {
		return err
	}

	defer func() { _ = client.Quit() }()

	auth := smtp.PlainAuth("", c.cfg.Username, c.cfg.Password, c.cfg.Host)
	err = client.Auth(auth)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) Dial() (*smtp.Client, error) {
	smtpServer := c.cfg.SMTPServer()
	encrypt := c.cfg.Encrypt
	if encrypt == "" {
		return smtp.Dial(smtpServer)
	}

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         c.cfg.Host,
	}

	switch encrypt {
	case "ssl":
		conn, err := tls.Dial("tcp", smtpServer, tlsConfig)
		if err != nil {
			return nil, err
		}

		return smtp.NewClient(conn, c.cfg.Host)

	case "tls":
		client, err := smtp.Dial(smtpServer)
		if err != nil {
			return nil, err
		}

		err = client.StartTLS(tlsConfig)
		return client, err

	default:
		return nil, errorpkg.ErrInternal(errors.New(fmt.Sprintf("Unknown encryption: %v", encrypt)))
	}
}

func (c *Client) SendMail(ctx context.Context, cmd *SendEmailCommand) error {
	if len(cmd.ToAddresses) == 0 {
		return errorpkg.ErrInternal(errors.New(fmt.Sprintf("Missing email address")))
	}

	addrs := make([]string, len(cmd.ToAddresses))
	for i, address := range cmd.ToAddresses {
		addrs[i] = address
	}

	err := c.sendMail(ctx, addrs, cmd)
	if err != nil {
		return errorpkg.ErrInternal(errors.New(fmt.Sprintf("Không thể gửi email đến địa chỉ %v (%v). Nếu cần thêm thông tin, vui lòng liên hệ %v.")))

	}
	return nil
}

func (c *Client) sendMail(ctx context.Context, addresses []string, cmd *SendEmailCommand) error {
	client, err := c.Dial()
	if err != nil {
		return err
	}
	defer func() { _ = client.Quit() }()

	auth := smtp.PlainAuth("", c.cfg.Username, c.cfg.Password, c.cfg.Host)
	err = client.Auth(auth)
	if err != nil {
		fmt.Errorf(err.Error())
		return err
	}

	//mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	//subject := "=?utf-8?B?" + base64.StdEncoding.EncodeToString([]byte(cmd.Subject)) + "?="
	return nil
	//var errs errorpkg.Errors
	//for _, email := range addresses {
	//	msg := []byte(fmt.Sprintf(
	//		"From: %s <%s> \r\nTo: %s\r\nSubject: %s\r\n%s\r\n\r\n%s\r\n",
	//		cmd.FromName, c.cfg.FromAddress, email, subject, mime, cmd.Content))
	//
	//	err = client.Mail(c.cfg.FromAddress)
	//	if err != nil {
	//		errs = append(errs, err)
	//		continue
	//	}
	//	err = client.Rcpt(email)
	//	if err != nil {
	//		errs = append(errs, err)
	//		continue
	//	}
	//	d, err := client.Data()
	//	if err != nil {
	//		errs = append(errs, err)
	//		continue
	//	}
	//	if _, err := d.Write(msg); err != nil {
	//		errs = append(errs, err)
	//		continue
	//	}
	//	err = d.Close()
	//	if err != nil {
	//		errs = append(errs, err)
	//		continue
	//	}
	//}
	//if len(errs) > 0 {
	//	fmt.Errorf("Can not send email", err)
	//}
	//return errs.Any()
}
