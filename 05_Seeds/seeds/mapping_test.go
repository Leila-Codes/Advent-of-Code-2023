package seeds

import "testing"

func TestMappingRange_Contains(t *testing.T) {
	mp := MappingRangeFromText("50 98 2")

	if mp.DestinationStart != 50 {
		t.Errorf("mapping range destination is not correct. expected '%d' but got '%d'", 50, mp.DestinationStart)
		t.Fail()
	}

	if mp.SourceStart != 98 {
		t.Errorf("mapping range source is not correct. expected '%d' but got '%d'", 98, mp.SourceStart)
		t.Fail()
	}

	if mp.Length != 2 {
		t.Errorf("mapping range length is not correct. expected '%d' but got '%d'", 2, mp.Length)
	}

	if !mp.Contains(98) {
		t.Error("mapping range did not contain 98!?")
		t.Fail()
	}

	if v := mp.Map(98); v != 50 {
		t.Errorf("mapping did not translate correctly. expected '%d' but got '%d'", 50, v)
		t.Fail()
	}

	if v := mp.Map(99); v != 51 {
		t.Errorf("mapping did not translate correctly. expected '%d' but got '%d'", 51, v)
		t.Fail()
	}

	// now try one not in the mapping
	if v := mp.Map(100); v != 100 {
		t.Errorf("mapping did not translate correctly. expected '%[1]d' but got '%[1]d'", 100)
		t.Fail()
	}
}
