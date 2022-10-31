package main

import (
	"fmt"

	"github.com/rawdaGastan/dmidecode/pkg"
)

func main() {
	dmiDecode := pkg.NewDMIDecoder()
	dmiDecode.Decode()

	//fmt.Printf("dmiDecode.GetSections(): %v\n", dmiDecode.GetSections())

	//fmt.Println(dmiDecode.GetSection("BIOS Information"))

	//fmt.Printf("dmiDecode.GetOptions(\"BIOS Information\"): %v\n", dmiDecode.GetOptions("BIOS Information"))

	op, err := dmiDecode.GetList("BIOS Information", "Characteristics")

	fmt.Printf("op: %v %v\n", op, err)

	op2, err := dmiDecode.Get("BIOS Information", "Vendor") // -> Dell Inc.

	fmt.Printf("op: %v %v\n", op2, err)

}
