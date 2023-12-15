package seeds

import "testing"

func TestMappingCollection_Map(t *testing.T) {
	mapping := MappingCollectionFromText(`seed-to-soil map:
50 98 2
52 50 48`)

	if len(mapping.Ranges) != 2 {
		t.Errorf("mapping doesn't have enough ranges. expected '%d' but got '%d'", 2, len(mapping.Ranges))
	}

	if v := mapping.Map(98); v != 50 {
		t.Errorf("mapping did not apply correctly. expected '%d' but got '%d'", 50, v)
		t.Fail()
	}

	if v := mapping.Map(50); v != 52 {
		t.Errorf("mapping did not apply correctly. expected '%d' but got '%d'", 52, v)
		t.Fail()
	}

	if v := mapping.Map(79); v != 81 {
		t.Errorf("mapping did not apply correctly. expected '%d' but got '%d'", 81, v)
	}
}
