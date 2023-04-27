package domain

type IEvent interface {
	GetEvent() (string, []byte)
}
