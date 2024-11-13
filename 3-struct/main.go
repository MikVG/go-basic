package main

import (
	"demo/struct/bins"
	"fmt"
)

func main() {
	myBin1 := bins.NewBin("1-bin", true, "Bin 1")
	myBin2 := bins.NewBin("2-bin", true, "Bin 2")
	myBin3 := bins.NewBin("3-bin", true, "Bin 3")

	myBinList := bins.NewBinList()

	myBinList.Add(myBin1)
	myBinList.Add(myBin2)
	myBinList.Add(myBin3)
	fmt.Println(*myBinList)
}
