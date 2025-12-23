package Gym

import "fmt"

type Member interface {
	GetDetails() string
}

type BasicMember struct {
	Name string
}

func (b BasicMember) GetDetails() string {
	return "Basic member: " + b.Name
}

type PremiumMember struct {
	Name string
}

func (p PremiumMember) GetDetails() string {
	return "Premium member: " + p.Name
}

type Gym struct {
	Members map[uint64]Member
}

func NewGym() Gym {
	return Gym{Members: make(map[uint64]Member)}
}

func (g *Gym) AddMember(id uint64, m Member) {
	g.Members[id] = m
}

func (g *Gym) ListMembers() {
	for _, m := range g.Members {
		fmt.Println(m.GetDetails())
	}
}
