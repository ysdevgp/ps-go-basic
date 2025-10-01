package main

import (
	"time"
)

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

type BinList struct {
	bins []Bin
}

func main() {

}

func NewBin(id string, private bool, name string) (*Bin, error) {
	return &Bin{
		id:        id,
		private:   private,
		createdAt: time.Now(),
		name:      name,
	}, nil
}

func NewBinList(length int) (*BinList, error) {
	return &BinList{
		bins: make([]Bin, length),
	}, nil
}
