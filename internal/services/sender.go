package services

type ISender interface {
	Send(text string) error
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

func (s *Sender) Send(text string) error {
	err := s.sender.Send(text)
	return err
}
