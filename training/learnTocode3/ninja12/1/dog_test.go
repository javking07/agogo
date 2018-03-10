package dog

import "testing"

func TestYears(t *testing.T) {
	v := Years(7)

	if v != 49 {
		t.Error("this aint dog years bruh")
	}
}
