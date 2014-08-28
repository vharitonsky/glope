package glope

import (
	"math"
	"testing"
)

func TestClusterStruct(t *testing.T) {
	trans := &Transaction{Instance: "test", Items: []string{"test", "transaction"}}
	trans2 := &Transaction{Instance: "test2", Items: []string{"test2", "transaction2"}}
	cluster := newCluster(0, trans)
	cluster.addItem("test item")
	cluster.removeItem("test item")
	cluster.addTransaction(trans2)
	cluster.getItemsProfit([]string{"test2", "transaction"}, 4.0)
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
