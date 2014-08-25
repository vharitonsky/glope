package glope

import (
	"fmt"
	"math"
)

type Cluster struct {
	id           int
	n            int            //Number of transactions
	w            int            //Number of unique items
	s            int            //Total number of items
	occ          map[string]int //Item to item count map
	Transactions []*Transaction
}

type Transaction struct {
	cluster         *Cluster
	clusterPosition int //Position of transaction inside cluster
	Instance        interface{}
	Items           []string
}

//Calculate a profit using this formula: number of items/(number of unique items ** repulsion)
func getProfit(s, w int, r float64) float64 {
	return float64(s) / math.Pow(float64(w), r)
}

func newCluster(id int, trans *Transaction) *Cluster {
	items_len := len(trans.Items)
	c := &Cluster{
		id:           id,
		occ:          make(map[string]int, items_len),
		s:            items_len,
		w:            items_len,
		n:            1,
		Transactions: []*Transaction{trans},
	}
	for _, item := range trans.Items {
		c.occ[item] = 1
	}
	trans.cluster = c
	return c
}

func (c *Cluster) String() string {
	return fmt.Sprintf("[Cluster %d]", c.id)
}

//Calculates a profit of adding transaction items to given cluster
func (c *Cluster) getItemsProfit(items []string, r float64) float64 {
	if c.n == 0 {
		s := len(items)
		return getProfit(s, s, r)
	} else {
		sNew := c.s + len(items)
		wNew := c.w
		for _, item := range items {
			if _, found := c.occ[item]; !found {
				wNew++
			}
		}
		return getProfit(sNew*(c.n+1), wNew, r) - getProfit(c.s*c.n, c.w, r)
	}
}

//Adds an item of a transaction to the cluster
func (c *Cluster) addItem(item string) {
	val, found := c.occ[item]
	if !found {
		c.occ[item] = 1
	} else {
		c.occ[item] = val + 1
	}
}

//Removes an item of a transaction from the cluster
func (c *Cluster) removeItem(item string) {
	val, found := c.occ[item]
	if !found {
		return
	}
	if val == 1 {
		delete(c.occ, item)
	}
	c.occ[item] -= 1
}

func (c *Cluster) addTransaction(trans *Transaction) {
	for _, item := range trans.Items {
		c.addItem(item)
	}
	c.s += len(trans.Items)
	c.w = len(c.occ)
	c.n++
	trans.clusterPosition = len(c.Transactions)
	c.Transactions = append(c.Transactions, trans)
	trans.cluster = c
}

func (c *Cluster) removeTransaction(trans *Transaction) {
	for _, item := range trans.Items {
		c.removeItem(item)
	}
	c.s -= len(trans.Items)
	c.w = len(c.occ)
	c.n--
	c.Transactions[trans.clusterPosition] = nil
}

func (c *Cluster) clearNilTransactions() {
	nonNilTransactions := make([]*Transaction, 0)
	for _, transaction := range c.Transactions {
		if transaction != nil {
			nonNilTransactions = append(nonNilTransactions, transaction)
		}
	}
	c.Transactions = nonNilTransactions
}

//Clusterizes given transactions
func Clusterize(data []*Transaction, repulsion float64) []*Cluster {
	if repulsion == 0 {
		repulsion = 4.0 // default value
	}
	var clusters []*Cluster
	for _, transaction := range data {
		clusters = addTransactionToBestCluster(clusters, transaction, repulsion)
	}
	for {
		moved := false
		for _, transaction := range data {
			originalClusterId := transaction.cluster.id
			transaction.cluster.removeTransaction(transaction)
			clusters = addTransactionToBestCluster(clusters, transaction, repulsion)
			if transaction.cluster.id != originalClusterId {
				moved = true
			}
		}
		if !moved {
			break
		}
	}
	notEmptyClusters := make([]*Cluster, 0)
	for _, cluster := range clusters {
		if cluster.n > 0 {
			cluster.clearNilTransactions()
			notEmptyClusters = append(notEmptyClusters, cluster)
		}
	}
	return notEmptyClusters
}

func addTransactionToBestCluster(clusters []*Cluster, transaction *Transaction, repulsion float64) []*Cluster {
	tempS := len(transaction.Items)
	profitMax := getProfit(tempS, tempS, repulsion)

	var bestCluster *Cluster
	var bestProfit float64

	for _, cluster := range clusters {
		clusterProfit := cluster.getItemsProfit(transaction.Items, repulsion)
		if clusterProfit > bestProfit {
			if clusterProfit > profitMax {
				cluster.addTransaction(transaction)
				return clusters
			} else {
				bestCluster = cluster
				bestProfit = clusterProfit
			}
		}
	}
	if bestProfit >= profitMax {
		bestCluster.addTransaction(transaction)
		return clusters
	}
	return append(clusters, newCluster(len(clusters), transaction))
}
