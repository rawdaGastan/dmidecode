package pkg

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
)

type Entry = map[string][]string
type DMIMap = map[string]Entry

type DMIDecoder struct {
	decodedMap DMIMap
}

// NewCheckerService creates new instance from the parser
func NewDMIDecoder() DMIDecoder {
	return DMIDecoder{
		decodedMap: DMIMap{},
	}
}

func (dmi *DMIDecoder) GetDMIDecodeOutput() (string, error) {
	dmidecode := exec.Command("sudo", "dmidecode")
	out, err := dmidecode.Output()

	if err != nil {
		return "", err
	}

	if len(out) == 0 {
		return "", fmt.Errorf("dmidecode is not installed in your device")
	}

	return string(out), nil
}

// Decodes the string output of dmidecode
func (dmi *DMIDecoder) Decode(output string) error {

	if output == "" {
		return fmt.Errorf("the given data is empty")
	}

	scanner := bufio.NewScanner(strings.NewReader(output))
	key := ""
	val := []string{}
	parent := ""
	isParent := false

	// skip first 4 lines
	for i := 0; i < 4; i++ {
		scanner.Scan()
	}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "Handle") {
			key = ""
			val = []string{}
			parent = ""
			isParent = true
			continue
		}

		if string(line[0]) != "\t" && string(line[1]) != "\t" {

			parent = line
			dmi.decodedMap[parent] = Entry{}

		} else if isParent {

			line = strings.TrimLeft(line, "\t")
			parent = line
			dmi.decodedMap[parent] = Entry{}

		} else if string(line[0]) == "\t" && string(line[1]) != "\t" {

			line = strings.TrimLeft(line, "\t")
			keyVal := strings.Split(line, ": ")

			if len(keyVal) == 2 {

				key = keyVal[0]
				val = []string{keyVal[1]}

			} else {

				key = strings.ReplaceAll(keyVal[0], ":", "")
				val = []string{}

			}
			dmi.decodedMap[parent][key] = val

		} else {
			line = strings.TrimLeft(line, "\t")

			val = append(val, line)
			dmi.decodedMap[parent][key] = val
		}

		isParent = false

	}

	return nil

}

// Get all possible keys of the decoded output
func (dmi *DMIDecoder) GetSections() []string {

	var sections []string

	for parentKey := range dmi.decodedMap {
		sections = append(sections, parentKey)
	}

	return sections
}

// GetSection returns a dictionary for the section given --> map[key1: val1 key2: val2 ....]
func (dmi *DMIDecoder) GetSection(sectionKey string) (Entry, error) {

	var section Entry

	if section, exists := dmi.decodedMap[sectionKey]; exists {
		return section, nil
	}

	return section, fmt.Errorf("the given section key %v does not exist", sectionKey)
}

// Get all possible options of a given key of the decoded output
func (dmi *DMIDecoder) GetOptions(sectionKey string) []string {

	section, _ := dmi.GetSection(sectionKey)
	var options []string

	for key := range section {
		options = append(options, key)
	}

	return options
}

// Get a string value of a given section key
func (dmi *DMIDecoder) Get(sectionKey string, optionKey string) (string, error) {

	option := ""

	if option, exists := dmi.decodedMap[sectionKey][optionKey]; exists {
		if len(option) == 1 {
			return option[0], nil
		} else {
			return "", fmt.Errorf("option '%v' exists as a list in the given key '%v'. you can use GetList()", optionKey, sectionKey)
		}
	}

	return option, fmt.Errorf("option '%v' does not exist in the given key '%v'", optionKey, sectionKey)
}

// Get a list value of a given section key
func (dmi *DMIDecoder) GetList(sectionKey string, optionKey string) ([]string, error) {

	option := []string{}

	if option, exists := dmi.decodedMap[sectionKey][optionKey]; exists {
		return option, nil
	}

	return option, fmt.Errorf("option '%v' does not exist in the given key '%v'", optionKey, sectionKey)
}
