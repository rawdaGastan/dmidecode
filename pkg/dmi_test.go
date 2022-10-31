package pkg

import (
	"testing"
)

func TestDMIDecode(t *testing.T) {

	t.Run("installed", func(t *testing.T) {
		dmi := NewDMIDecoder()
		_, err := dmi.GetDMIDecodeOutput()

		if err != nil {
			t.Errorf("dmidecode should be installed")
		}
	})

	t.Run("valid_output", func(t *testing.T) {
		dmi := NewDMIDecoder()
		out, _ := dmi.GetDMIDecodeOutput()

		if out == "" {
			t.Errorf("dmidecode output should be returned")
		}
	})

	t.Run("valid_decoding", func(t *testing.T) {
		dmi := NewDMIDecoder()
		out, _ := dmi.GetDMIDecodeOutput()
		err := dmi.Decode(out)

		if err != nil {
			t.Errorf("the given data is valid")
		}
	})

	t.Run("invalid_decoding", func(t *testing.T) {
		dmi := NewDMIDecoder()
		err := dmi.Decode("")

		if err == nil {
			t.Errorf("the given data is invalid")
		}
	})
}

func TestDMIFunctions(t *testing.T) {

	t.Run("test_sections", func(t *testing.T) {
		dmi := NewDMIDecoder()
		out, _ := dmi.GetDMIDecodeOutput()
		err := dmi.Decode(out)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		sections := dmi.GetSections()

		if len(sections) == 0 {
			t.Errorf("sections should not be empty")
		}
	})

	t.Run("test_empty_sections", func(t *testing.T) {
		dmi := NewDMIDecoder()
		err := dmi.Decode("empty")

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		sections := dmi.GetSections()

		if len(sections) > 0 {
			t.Errorf("sections should be empty")
		}
	})

	t.Run("test_options", func(t *testing.T) {
		dmi := NewDMIDecoder()
		out, _ := dmi.GetDMIDecodeOutput()
		err := dmi.Decode(out)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		options := dmi.GetOptions("BIOS Information")

		if len(options) == 0 {
			t.Errorf("options of 'BIOS Information' should not be empty")
		}
	})

	t.Run("test_empty_options", func(t *testing.T) {
		dmi := NewDMIDecoder()
		out, _ := dmi.GetDMIDecodeOutput()
		err := dmi.Decode(out)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		options := dmi.GetOptions("Information")

		if len(options) > 0 {
			t.Errorf("options of 'Information' should be empty")
		}
	})

	t.Run("test_section", func(t *testing.T) {
		dmi := NewDMIDecoder()
		out, _ := dmi.GetDMIDecodeOutput()
		err := dmi.Decode(out)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		section, _ := dmi.GetSection("BIOS Information")

		if len(section) == 0 {
			t.Errorf("section of 'BIOS Information' should not be empty")
		}
	})

	t.Run("test_wrong_section", func(t *testing.T) {
		dmi := NewDMIDecoder()
		out, _ := dmi.GetDMIDecodeOutput()
		errDecode := dmi.Decode(out)

		if errDecode != nil {
			t.Errorf("unexpected error: %v", errDecode)
		}

		_, err := dmi.GetSection("BIOSS Information")

		if err == nil {
			t.Errorf("section of 'BIOSS Information' should not exist")
		}
	})

}

func TestDMIGetters(t *testing.T) {

	t.Run("test_get_value", func(t *testing.T) {
		dmi := NewDMIDecoder()
		out, _ := dmi.GetDMIDecodeOutput()
		err := dmi.Decode(out)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		value, _ := dmi.Get("BIOS Information", "Vendor")

		if value != "Dell Inc." {
			t.Errorf("value should be Dell Inc.")
		}
	})

	t.Run("test_get_list_value", func(t *testing.T) {
		dmi := NewDMIDecoder()
		out, _ := dmi.GetDMIDecodeOutput()
		errDecode := dmi.Decode(out)

		if errDecode != nil {
			t.Errorf("unexpected error: %v", errDecode)
		}

		_, err := dmi.Get("BIOS Information", "Characteristics")

		if err == nil {
			t.Errorf("'Characteristics' has a list value")
		}
	})

	t.Run("test_get_string_value", func(t *testing.T) {
		dmi := NewDMIDecoder()
		out, _ := dmi.GetDMIDecodeOutput()
		errDecode := dmi.Decode(out)

		if errDecode != nil {
			t.Errorf("unexpected error: %v", errDecode)
		}

		_, err := dmi.Get("BIOS Information", "Vendor")

		if err != nil {
			t.Errorf("'Vendor' has a string value")
		}
	})

	t.Run("test_get_wrong_key", func(t *testing.T) {
		dmi := NewDMIDecoder()
		out, _ := dmi.GetDMIDecodeOutput()
		errDecode := dmi.Decode(out)

		if errDecode != nil {
			t.Errorf("unexpected error: %v", errDecode)
		}

		_, err := dmi.Get("BIOS Information", " Vendor")

		if err == nil {
			t.Errorf("' Vendor' should not exist")
		}
	})

	t.Run("test_list_multi_value", func(t *testing.T) {
		dmi := NewDMIDecoder()
		out, _ := dmi.GetDMIDecodeOutput()
		err := dmi.Decode(out)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		value, _ := dmi.GetList("BIOS Information", "Characteristics")

		if len(value) == 1 {
			t.Errorf("'Characteristics' should have a length greater than 1")
		}
	})

	t.Run("test_list_single_value", func(t *testing.T) {
		dmi := NewDMIDecoder()
		out, _ := dmi.GetDMIDecodeOutput()
		err := dmi.Decode(out)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		value, _ := dmi.GetList("BIOS Information", "Vendor")

		if len(value) != 1 {
			t.Errorf("'Vendor' should have a length equals 1")
		}
	})

	t.Run("test_list_wrong_key", func(t *testing.T) {
		dmi := NewDMIDecoder()
		out, _ := dmi.GetDMIDecodeOutput()
		errDecode := dmi.Decode(out)

		if errDecode != nil {
			t.Errorf("unexpected error: %v", errDecode)
		}

		_, err := dmi.GetList("BIOS Information", " Vendor")

		if err == nil {
			t.Errorf("' Vendor' should not exist")
		}
	})
}
