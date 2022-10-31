/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"fmt"

	"github.com/rawdaGastan/dmidecode/pkg"
)

func main() {
	dmiDecode := pkg.NewDMIDecoder()
	dmiDecode.Decode()

	fmt.Printf("dmiDecode.GetSections(): %v\n", dmiDecode.GetSections())

	//fmt.Println(dmiDecode.GetSection("BIOS Information"))

	//fmt.Printf("dmiDecode.GetOptions(\"BIOS Information\"): %v\n", dmiDecode.GetOptions("BIOS Information"))

	op, err := dmiDecode.Get("BIOS Information", "Characteristics")

	fmt.Printf("op: %v %v\n", op, err)

	op, err = dmiDecode.Get("BIOS Information", "Vendor") // -> Dell Inc.

	fmt.Printf("op: %v %v\n", op, err)

}
