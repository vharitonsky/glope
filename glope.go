package glope

import (
	"fmt"
	"log"
	"math"
)

type Cluster struct {
	id           int
	n            float64        //Number of transactions
	w            float64        //Number of unique items
	s            float64        //Total number of items
	occ          map[string]int //Item to item count map
	Transactions []*Transaction
}

type Transaction struct {
	cluster         *Cluster
	clusterPosition int //Position of transaction inside cluster
	Instance        interface{}
	Items           []string
}

func getProfit(s, w, r float64) float64 {
	return s / math.Pow(w, r)
}

func newCluster(id int) *Cluster {
	return &Cluster{id: id, n: 0, w: 0, s: 0, Transactions: make([]*Transaction, 0), occ: make(map[string]int, 0)}
}

func (c *Cluster) String() string {
	return fmt.Sprintf("[Cluster %d]", c.id)
}

func (c *Cluster) getProfit(items []string, r float64) float64 {
	sNew := c.s + float64(len(items))
	wNew := c.w
	for _, item := range items {
		if _, found := c.occ[item]; !found {
			wNew++
		}
	}
	if c.n == 0 {
		return getProfit(sNew, wNew, r)
	} else {
		profit := getProfit(c.s*c.n, c.w, r)
		profitNew := getProfit(sNew*(c.n+1), wNew, r)
		return profitNew - profit
	}
}

func (c *Cluster) addItem(item string) {
	val, found := c.occ[item]
	if !found {
		c.occ[item] = 1
	} else {
		c.occ[item] = val + 1
	}
	c.s++
}

func (c *Cluster) removeItem(item string) {
	val, found := c.occ[item]
	if !found {
		return
	}
	if val == 1 {
		delete(c.occ, item)
	}
	c.occ[item] -= 1
	c.s--
}

func (c *Cluster) addTransaction(trans *Transaction) {
	for _, item := range trans.Items {
		c.addItem(item)
	}
	c.w = float64(len(c.occ))
	c.n++
	trans.clusterPosition = len(c.Transactions)
	c.Transactions = append(c.Transactions, trans)
	trans.cluster = c
}

func (c *Cluster) removeTransaction(trans *Transaction) {
	for _, item := range trans.Items {
		c.removeItem(item)
	}
	c.w = float64(len(c.occ))
	c.n--
	trans.cluster = nil
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

func Clusterize(data []*Transaction, repulsion float64) []*Cluster {
	if repulsion == 0 {
		repulsion = 4.0 // default value
	}
	var clusters []*Cluster
	log.Print("Initializing clusters")
	for _, transaction := range data {
		clusters = addTransactionToBestCluster(clusters, transaction, repulsion)
	}
	log.Printf("Init finished, created %d clusters", len(clusters))
	log.Print("Moving transactions to best clusters")
	for i := 1; ; i++ {
		log.Printf("move %d", i)
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
	log.Print("Finished, cleaning empty clusters")
	notEmptyClusters := make([]*Cluster, 0)
	for _, cluster := range clusters {
		if cluster.n > 0 {
			cluster.clearNilTransactions()
			notEmptyClusters = append(notEmptyClusters, cluster)
		}
	}
	log.Printf("Cleaning finished, returning %d clusters", len(notEmptyClusters))
	return notEmptyClusters
}

func addTransactionToBestCluster(clusters []*Cluster, transaction *Transaction, repulsion float64) []*Cluster {
	if len(clusters) > 0 {
		tempS := float64(len(transaction.Items))
		tempW := tempS
		profitMax := getProfit(tempS, tempW, repulsion)

		var bestCluster *Cluster
		var bestProfit float64

		for _, cluster := range clusters {
			clusterProfit := cluster.getProfit(transaction.Items, repulsion)
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
	}

	cluster := newCluster(len(clusters))
	cluster.addTransaction(transaction)
	return append(clusters, cluster)
}
