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
	if !cluster.hasItem("test item") {
		t.Error("'test item' was supposed to be inside cluster")
		t.FailNow()
	}
	cluster.removeItem("test item")
	if cluster.hasItem("test item") {
		t.Error("'test item' was supposed to be removed from cluster")
		t.FailNow()
	}

	cluster.addTransaction(trans2)
	if !cluster.hasTransaction(trans2) {
		t.Error("Transaction ", trans.Instance, " was supposed to be added to cluster")
		t.FailNow()
	}
	cluster.removeTransaction(trans2)
	if cluster.hasTransaction(trans2) {
		t.Error("Transaction ", trans.Instance, " was supposed to be removed from cluster")
		t.FailNow()
	}
	cluster.removeTransaction(trans)

	expectedProfit := getProfit(2, 2, 4.0)
	calculatedProfit := cluster.getItemsProfit([]string{"test2", "transaction"}, 4.0)

	if calculatedProfit != expectedProfit {
		t.Error("Profit of adding transaction to empty cluster should be ", expectedProfit, " got ", calculatedProfit)
		t.FailNow()
	}

	addTransactionToBestCluster([]*Cluster{cluster}, trans, 4.0)

	if !cluster.hasTransaction(trans) {
		t.Error("Transaction was supposed to be added to an empty cluster")
		t.FailNow()
	}

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
