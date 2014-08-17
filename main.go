package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"math"
	"os"
	"strings"
)

var (
	input_file = flag.String("input", "", "Input file with transactions")
	output_dir = flag.String("output", "", "A dir to store created clusters to")
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

func newCluster(id int) *Cluster {
	return &Cluster{id: id, n: 0, w: 0, s: 0, instances: make(map[string]bool, 0), occ: make(map[string]int, 0)}
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
	log.Print(data)
	if repulsion == 0 {
		repulsion = 4.0 // default value
	}
	var clusters []*Cluster
	for _, transaction := range data {
		clusters = addTransactionToBestCluster(clusters, transaction, repulsion)
	}
	log.Print(clusters)
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
	log.Printf("Finished %v", clusters)
	return clusters
}

func addTransactionToBestCluster(clusters []*Cluster, transaction *Transaction, repulsion float64) []*Cluster {
	if len(clusters) > 0 {
		tempS := float64(len(transaction.items))
		tempW := tempS
		profitMax := getProfit(tempS, tempW, repulsion)
		log.Printf("Profit max %f", profitMax)
		for _, cluster := range clusters {
			clusterProfit := cluster.getProfit(transaction.items, repulsion)
			log.Printf("Cluster profit %f", clusterProfit)
			if clusterProfit >= profitMax {
				cluster.addTransaction(transaction)
				return clusters
			}
		}
	}

	cluster := newCluster(len(clusters))
	cluster.addTransaction(transaction)
	return append(clusters, cluster)
}

func main() {
	flag.Parse()
	if *input_file == "" {
		log.Fatal("You must provide input file")
	}
	// if output_dir == ""{
	// 	log.Fail("You must provide output dir")
	// }
	file, err := os.Open(*input_file)
	if err != nil {
		log.Fatalf("Cannot open config file at [%s]: [%s]\n", *input_file, err)
	}
	defer file.Close()
	r := bufio.NewReader(file)
	var transactions []*Transaction
	for {
		line, err := r.ReadString('\n')

		if err != nil && line == "" {
			if err == io.EOF {
				break
			}
			log.Fatalf("Error when reading file [%s]: [%s]\n", *input_file, err)
		}
		instance := strings.TrimSuffix(line, "\n")
		items := make([]string, 0)
		visited := make(map[string]bool)
		for _, item := range strings.Split(instance, " ") {
			if _, found := visited[item]; !found {
				items = append(items, item)
				visited[item] = true
			}
		}
		transactions = append(transactions, &Transaction{instance: instance, items: items})
	}
	clusterize(transactions, 4.0)

}
