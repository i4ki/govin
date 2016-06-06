package govin

import "testing"

func TestWMIGeneration(t *testing.T) {
	if wmiMap == nil {
		t.Errorf("wmiMap wasn't initialized")
		return
	}

	_, err := genWMIFromRanges(nil)

	if err == nil {
		t.Error("Must fail")
		return
	}

	wmiMap, err := genWMIFromRanges([]wmiRange{
		wmiRange{
			'A', 'A', 'A', 'H', "Africa", "South Africa",
		},
	})

	if err != nil {
		t.Error(err)
		return
	}

	for i := 'A'; i <= 'A'; i++ {
		for j := 'A'; j <= 'H'; j++ {
			sym := string([]rune{i, j})

			val, ok := wmiMap[sym]

			if !ok {
				t.Errorf("'%s' not found", sym)
				return
			}

			if val.continent != "Africa" {
				t.Errorf("Invalid continent = %s", val.continent)
				return
			}

			if val.country != "South Africa" {
				t.Errorf("Invalid country '%s'", val.country)
				return
			}
		}
	}

	wmiMap, err = genWMIFromRanges([]wmiRange{
		wmiRange{
			'A', 'J', 'A', 'N', "Africa", "Ivory Coast",
		},
	})

	if err != nil {
		t.Error(err)
		return
	}

	for i := 'A'; i <= 'A'; i++ {
		for j := 'J'; j <= 'N'; j++ {
			sym := string([]rune{i, j})

			val, ok := wmiMap[sym]

			if !ok {
				t.Errorf("'%s' not found", sym)
				return
			}

			if val.continent != "Africa" {
				t.Errorf("Invalid continent = %s", val.continent)
				return
			}

			if val.country != "Ivory Coast" {
				t.Errorf("Invalid country '%s'", val.country)
				return
			}
		}
	}

	_, err = genWMIMap()

	if err != nil {
		t.Error(err)
		return
	}
}
