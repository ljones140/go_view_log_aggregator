package main

import "net"

type Page struct {
	Name   string
	Visits int
}

func (p *Page) AddVisit(address net.IP) {
	p.Visits = p.Visits + 1
}
