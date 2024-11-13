package bins

import (
	"time"
)

type Bin struct {
	id       string
	private  bool
	createAt time.Time
	name     string
}

type BinList []Bin

func (binList *BinList) Add(bin *Bin) {
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
