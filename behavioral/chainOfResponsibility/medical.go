package main

import "fmt"

type medical struct {
	next department
}

func (m *medical) execute(p *patient) {
	if p.medicalDone {
		fmt.Println("Medicine already given to patient")
		m.next.execute(p)
		return
	}
	fmt.Println("Medical giving medicine to patient")
	p.medicalDone = true
	m.next.execute(p)
}

func (m *medical) setNext(next department) {
	m.next = next
}
