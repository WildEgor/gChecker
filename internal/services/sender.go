package services

type ISender interface {
	Send(to string, text string) error
}

type SenderData struct {
	to   string
	text string
}

type Sender struct {
	sender ISender
}

func InitSender(s ISender) *Sender {
	return &Sender{
		sender: s,
	}
}

func (s *Sender) SetTransport(t ISender) *Sender {
	s.sender = t
	return s
}

func (s *Sender) Send(data SenderData) error {
	err := s.sender.Send(data.to, data.text)
	return err
}
