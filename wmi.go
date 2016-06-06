package govin

import "errors"

type (
	wmiRange struct {
		low1, low2         rune
		high1, high2       rune
		continent, country string
	}

	WMI struct {
		code, continent, country string
	}

	WMIMap map[string]WMI
)

var (
	codes = []rune{
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N',
		'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'X', 'Y', 'Z', '1', '2', '3',
		'4', '5', '6', '7', '8', '9', '0',
	}

	africaRanges = []wmiRange{
		wmiRange{
			'A', 'A', 'A', 'H', "Africa", "South Africa",
		},
		wmiRange{
			'A', 'J', 'A', 'N', "Africa", "Ivory Coast",
		},
		wmiRange{
			'A', 'P', 'A', '0', "Africa", "not assigned",
		},
		wmiRange{
			'B', 'A', 'B', 'E', "Africa", "Angola",
		},
		wmiRange{
			'B', 'F', 'B', 'K', "Africa", "Kenya",
		},
		wmiRange{
			'B', 'L', 'B', 'R', "Africa", "Tanzania",
		},
		wmiRange{
			'B', 'S', 'B', '0', "Africa", "not assigned",
		},
		wmiRange{
			'C', 'A', 'C', 'E', "Africa", "Benin",
		},
		wmiRange{
			'C', 'F', 'C', 'K', "Africa", "Madagascar",
		},
		wmiRange{
			'C', 'L', 'C', 'R', "Africa", "Tunisia",
		},
		wmiRange{
			'C', 'S', 'C', '0', "Africa", "not assigned",
		},
		wmiRange{
			'D', 'A', 'D', 'E', "Africa", "Egypt",
		},
		wmiRange{
			'D', 'F', 'D', 'K', "Africa", "Morocco",
		},
		wmiRange{
			'D', 'L', 'D', 'R', "Africa", "Zambia",
		},
		wmiRange{
			'D', 'S', 'D', '0', "Africa", "not assigned",
		},
		wmiRange{
			'E', 'A', 'E', 'E', "Africa", "Ethiopia",
		},
		wmiRange{
			'E', 'F', 'E', 'K', "Africa", "Mozambique",
		},
		wmiRange{
			'E', 'L', 'E', '0', "Africa", "not assigned",
		},
		wmiRange{
			'F', 'A', 'F', 'E', "Africa", "Ghana",
		},
		wmiRange{
			'F', 'F', 'F', 'K', "Africa", "Nigeria",
		},
		wmiRange{
			'F', 'L', 'F', '0', "Africa", "not assigned",
		},
		wmiRange{
			'G', 'A', 'G', '0', "Africa", "not assigned",
		},
		wmiRange{
			'H', 'A', 'H', '0', "Africa", "not assigned",
		},
	}

	samericaRanges = []wmiRange{
		wmiRange{
			'8', 'A', '8', 'E', "South America", "Argentina",
		},
		wmiRange{
			'8', 'F', '8', 'K', "South America", "Chile",
		},
		wmiRange{
			'8', 'L', '8', 'R', "South America", "Ecuador",
		},
		wmiRange{
			'8', 'S', '8', 'W', "South America", "Peru",
		},
		wmiRange{
			'8', 'X', '8', '2', "South America", "Venezuela",
		},
		wmiRange{
			'8', '3', '8', '0', "South America", "not assigned",
		},
		wmiRange{
			'9', 'A', '9', 'E', "South America", "Brazil",
		},
		wmiRange{
			'9', 'F', '9', 'K', "South America", "Colombia",
		},
		wmiRange{
			'9', 'L', '9', 'R', "South America", "Paraguay",
		},
		wmiRange{
			'9', 'S', '9', 'W', "South America", "Uruguay",
		},
		wmiRange{
			'9', 'X', '9', '2', "South America", "Trinidad & Tobago",
		},
		wmiRange{
			'9', '3', '9', '9', "South America", "Brazil",
		},
		wmiRange{
			'9', '0', '9', '0', "South America", "not assigned",
		},
	}

	countryRanges = append(africaRanges, samericaRanges...)

	wmiMap WMIMap
)

func init() {
	var err error

	wmiMap, err = genWMIMap()

	if err != nil {
		panic(err)
	}
}

func toCode(r rune) int {
	for i := 0; i < len(codes); i++ {
		if codes[i] == r {
			return i
		}
	}

	return -1
}

func genWMIMap() (WMIMap, error) {
	return genWMIFromRanges(countryRanges)
}

func genWMIFromRanges(cranges []wmiRange) (WMIMap, error) {
	wmiMap := make(WMIMap)

	if len(cranges) == 0 {
		return nil, errors.New("No range")
	}

	for r := 0; r < len(cranges); r++ {
		rng := cranges[r]

		for i := toCode(rng.low1); i <= toCode(rng.high1); i++ {
			for j := toCode(rng.low2); j <= toCode(rng.high2); j++ {
				wmi := WMI{}
				ci := codes[i]
				cj := codes[j]

				wmi.code = string([]rune{ci, cj})

				wmi.continent = rng.continent
				wmi.country = rng.country

				wmiMap[wmi.code] = wmi
			}
		}
	}

	return wmiMap, nil
}

func GetWMIMap() WMIMap { return wmiMap }
