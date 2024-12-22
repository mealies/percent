package percent

import "testing"

func TestPercent(t *testing.T) {
	pct, total := 25, 100
	part := Percent(pct, total)

	if part != 25.0 {
		t.Fatalf("%d is wrong number for %d percent", int(part), pct)
	}

	pct, total = 50, 1000000
	part = Percent(pct, total)

	if part != 500000.0 {
		t.Fatalf("%d is wrong number for %d percent", int(part), pct)
	}
}
