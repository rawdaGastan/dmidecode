package pkg

/*
dmidecode["BIOS Information"]["Version"]


dmioutput = dmi.GetDMIOUtput()
dmi = NewDMIDecoder(dmioutput)

dmi.Get("BIOS Information"," Vendor") // -> Dell Inc.
dmi.Get("BIOS Information", "Characterstics")


func GetDMIOutput() string {

	dmidecode installed or not
	if installed {
		execute command and return the output
	}

}

type DMIDecoder struct {
	text string
}



func (dmidecoder* DMIDecoder) NewDMIDecoder(output){


}
func ... parse() {

}

*/
