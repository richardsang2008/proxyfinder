package model

import "fmt"

type ProxyLevel int

const (
	Anoymous ProxyLevel = iota
	Eliteproxy
	Transparent
	Unknown
)

type ProxyType int

const (
	HTTP ProxyType = iota
	SOCKET5
	SOCKET4
)

type ProxyRecord struct {
	IP                string
	Port              int
	CountryCode       string
	Country           string
	Type              ProxyLevel
	ProxyType         ProxyType
	IsHttps           bool
	IsCreateAccountOk bool
}

func (m ProxyRecord) ToString() string {
	return fmt.Sprintf("%s:%v", m.IP, m.Port)
}
