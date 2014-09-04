package main

import (
	"fmt"
	"github.com/vharitonsky/glope"
)

func main() {
	transactions := []*glope.Transaction{
		{
			//You can store anything here, an id or name to identify your transaction
			Instance: "Brandname1 Productname1 Series1 red",
			//Items is a series of unique tokens inside a transaction
			Items: []string{"Brandname1", "Productname1", "Series1", "red"},
		},
		{
			Instance: "Brandname1 Productname1 Series1 refurbished",
			Items:    []string{"Brandname1", "Productname1", "Series1", "refurbished"},
		},
		{
			Instance: "Brandname1 Productname1 Series1 blue",
			Items:    []string{"Brandname1", "Productname1", "Seris1", "blue"},
		},
		{
			Instance: "Brandname2 Productname2 Series2 black",
			Items:    []string{"Brandname2", "Productname2", "Series2", "black"},
		},
	}
	/*
		Clusterize takes two parameters: list of pointers to transactions and repulsion.
		Repulsion parameter is a threshold of likeness for transactions,
		the higher, the more precise. In real life, repulsion depends on
		the amount of items in your transaction, the more items you have,
		the higher the repulsion should be.
		But it is up to you to conclude by trial and error the value of
		repulsion parameter. Default is 4.
	*/
	clusters := glope.Clusterize(transactions, 1.5)
	for _, cluster := range clusters {
		for _, transaction := range cluster.Transactions {
			fmt.Printf("%s in %v\n", transaction.Instance, cluster)
		}
	}
	/*
		Brandname1 Productname1 Series1 red in [Cluster 0]
		Brandname1 Productname1 Series1 blue in [Cluster 0]
		Brandname1 Productname1 Series1 refurbished in [Cluster 0]
		Brandname2 Productname2 Series2 black in [Cluster 1]
	*/
}
