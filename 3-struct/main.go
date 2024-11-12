package main

import (
	"fmt"
	"time"
)

type Bin struct {
	id       string
	private  bool
	createAt time.Time
	name     string
}

type BinList []Bin

func (binList *BinList) add(bin *Bin) {
	*binList = append(*binList, *bin)
}

func NewBin(id string, private bool, name string) *Bin {
	return &Bin{
		id:       id,
		private:  private,
		createAt: time.Now(),
		name:     name,
	}
}

func NewBinList() *BinList {
	BinList := make(BinList, 0)
	return &BinList
}

func main() {
	myBin1 := NewBin("1-bin", true, "Bin 1")
	myBin2 := NewBin("2-bin", true, "Bin 2")
	myBin3 := NewBin("3-bin", true, "Bin 3")

	myBinList := NewBinList()

	myBinList.add(myBin1)
	myBinList.add(myBin2)
	myBinList.add(myBin3)
	fmt.Println(*myBinList)
}
