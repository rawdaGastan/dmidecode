package main

import (
	"fmt"

	"github.com/rawdaGastan/dmidecode/pkg"
)

func main() {
	dmiDecode := pkg.NewDMIDecoder()
	out, err := dmiDecode.GetDMIDecodeOutput()
	fmt.Printf("err: %v\n", err)

	err = dmiDecode.Decode(out)
	fmt.Printf("err: %v\n", err)

	//fmt.Printf("dmiDecode.GetSections(): %v\n", dmiDecode.GetSections())

	//fmt.Println(dmiDecode.GetSection("BIOS Information"))

	//fmt.Printf("dmiDecode.GetOptions(\"BIOS Information\"): %v\n", dmiDecode.GetOptions("BIOS Information"))

	op, err := dmiDecode.GetList("BIOS Information", "Characteristics")

	fmt.Printf("op: %v %v\n", op, err)

	op2, err := dmiDecode.Get("BIOS Information", "Vendor") // -> Dell Inc.

	fmt.Printf("op: %v %v\n", op2, err)

}
