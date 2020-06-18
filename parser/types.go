package parser

import (
	"net/url"
)

type Player struct {
	Platform   string
	Name       string
	Rank       Rank
	url        url.URL
	Endorsment Endorsment
	Heroes     map[string][]Stat
}

type Rank struct {
	DD   int
	Tank int
	Heal int
}

type Endorsment struct {
	Level         int
	Shotcaller    float64
	Teammate      float64
	Sportsmanship float64
}

type Stat struct {
	Name  string
	Value SValue
	Help  string
}

type SValue struct {
	QP          float64
	Competitive float64
}
