package main

import (
	"math/rand"
	"time"
)

type Quote struct {
	ID     int    `json:"id"`
	Text   string `json:"text"`
	Author string `json:"author"`
}

var Quotes []Quote

var Rng *rand.Rand

func init() {
	Rng = rand.New(rand.NewSource(time.Now().UnixNano()))

	Quotes = []Quote{
		{ID: 1, Text: "The best way to get started is to quit talking and begin doing.", Author: "Walt Disney"},
		{ID: 2, Text: "Don’t let yesterday take up too much of today.", Author: "Will Rogers"},
		{ID: 3, Text: "It’s not whether you get knocked down, it’s whether you get up.", Author: "Vince Lombardi"},
		{ID: 4, Text: "Whether you think you can or you think you can’t, you’re right.", Author: "Henry Ford"},
	}
}
