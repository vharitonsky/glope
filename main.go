package main

import (
	"math"
)

type Cluster struct {
	id        int
	n         float64
	w         float64
	s         float64
	instances map[string]bool
	occ       map[string]int
}

type Transaction struct {
	cluster  *Cluster
	instance string
	items    []string
}

func getProfit(s, w, r float64) float64 {
	return s / math.Pow(w, r)
}

func (c *Cluster) addItem(item string) {
	val, found := c.occ[item]
	if !found {
		c.occ[item] = 1
	} else {
		c.occ[item] = val + 1
	}
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
	for _, item := range trans.items {
		c.addItem(item)
	}
	c.w = float64(len(c.occ))
	c.n++
	c.instances[trans.instance] = true
	trans.cluster = c
}

func (c *Cluster) removeTransaction(trans *Transaction) {
	for _, item := range trans.items {
		c.removeItem(item)
	}
	c.w = float64(len(c.occ))
	c.n--
	delete(c.instances, trans.instance)
	trans.cluster = nil
}

func clusterize(data []*Transaction, repulsion float64) []*Cluster {
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
	return clusters
}

func addTransactionToBestCluster(clusters []*Cluster, transaction *Transaction, repulsion float64) []*Cluster {
	if len(clusters) > 0 {
		tempS := float64(len(transaction.items))
		tempW := tempS
		profitMax := getProfit(tempS, tempW, repulsion)
		for _, cluster := range clusters {
			if cluster.getProfit(transaction.items, repulsion) > profitMax {
				cluster.addTransaction(transaction)
				return clusters
			}
		}
	}

	cluster := Cluster{id: len(clusters)}
	cluster.addTransaction(transaction)
	return append(clusters, &cluster)
}

func main() {

}
