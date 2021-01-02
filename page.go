package main

import "net"

type Page struct {
	Name        string
	Visits      int
	UniqueViews int
	vistorIPs   []net.IP
}

func (p *Page) AddVisit(visitIP net.IP) {
	p.Visits++
	if p.visitAddressUnique(visitIP) {
		p.UniqueViews++
		p.vistorIPs = append(p.vistorIPs, visitIP)
	}
}

func (p *Page) visitAddressUnique(visitIP net.IP) bool {
	for _, ip := range p.vistorIPs {
		if ip.Equal(visitIP) {
			return false
		}
	}
	return true
}
