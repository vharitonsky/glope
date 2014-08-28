package glope

import (
	"math"
	"testing"
)

func TestClusterStruct(t *testing.T) {
	//stub
}

func TestGetProfit(t *testing.T) {
	s := 5
	w := 3
	r := 4.0
	expectedProfit := float64(s) / math.Pow(float64(w), r)
	calculatedProfit := getProfit(s, w, r)
	if expectedProfit != calculatedProfit {
		t.Error("Expected profit was ", expectedProfit, " got ", calculatedProfit)
		t.FailNow()
	}
}
