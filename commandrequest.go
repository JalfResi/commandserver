package commandserver

import (
	"net/textproto"
)

type CommandRequest struct {
	Method  string
	Headers textproto.MIMEHeader
	Body    string
}

func NewCommandRequest(c *textproto.Conn) (*CommandRequest, error) {
	verb, err := c.ReadLine()
	if err != nil {
		return nil, err
	}

	headers, err := c.ReadMIMEHeader()
	if err != nil {
		return nil, err
	}

	// read body here

	return &CommandRequest{
		Method:  verb,
		Headers: headers,
		//Body: body,
	}, nil
}
