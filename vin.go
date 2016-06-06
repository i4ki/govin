package govin

import "fmt"

type (
	VIN struct {
		wmi WMI // Wold Manufacturer Identifier
		//		vds VDS // Vehicle Descriptor Section
		//		vis VIS // Vehicle Identification Section
	}
)

func Decode(vinStr string) (*VIN, error) {
	if len(vinStr) != 17 {
		return nil, fmt.Errorf("Wrong VIN length. Correct is 17, but given was '%d'", len(vinStr))
	}

	var (
		vin          VIN
		wmi          WMI
		manufacturer WMIManufacturer
		ok           bool
	)

	wmiStr := vinStr[0:3]

	if wmi, ok = wmiMap[wmiStr[0:2]]; !ok {
		return nil, fmt.Errorf("Invalid WMI code '%s'", wmiStr)
	}

	if manufacturer, ok = manufacturerMap[wmiStr]; !ok {
		wmi.manufacturer = WMIManufacturer{
			code: wmiStr,
			name: "Unknown",
		}
	} else {
		wmi.manufacturer = manufacturer
	}

	wmi.code = manufacturer.code
	vin.wmi = wmi
	return &vin, nil
}
