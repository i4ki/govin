package govin

import (
	"fmt"
	"testing"
)

func TestVINDecode(t *testing.T) {
	vinStr := "93HES15504Z106171"

	_, err := Decode("A")

	if err == nil {
		t.Errorf("Must fail")
		return
	}

	vin, err := Decode(vinStr)

	if err != nil {
		t.Error(err)
		return
	}

	if vin.wmi.code != "93H" {
		t.Errorf("Invalid WMI code: %s", vin.wmi.code)
		return
	}

	if vin.wmi.continent != "South America" ||
		vin.wmi.country != "Brazil" ||
		vin.wmi.manufacturer.code != "93H" ||
		vin.wmi.manufacturer.name != "Honda Brazil" {
		t.Errorf("Invalid parsed VIN: %+v\n", vin)
		return
	}

	fmt.Printf("Region: %s\n", vin.wmi.continent)
	fmt.Printf("Country: %s\n", vin.wmi.country)
	fmt.Printf("Manufacturer: %s\n", vin.wmi.manufacturer.name)
	fmt.Printf("Manufacturer code: %s\n", vin.wmi.manufacturer.code)
}
