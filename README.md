# Glope - a Go implementation of Clope algorythm of clusterization.

What is this?
=========

A library that implements Clope algorythm of clusterization.
You may need it to clusterize an array of transactions with overlapping items.
Like, for example, names of products that differ by minor tokens.

Example:
======

```
Brandname1 Productname1 Series1 red
Brandname1 Productname1 Series1 blue
Brandname1 Productname1 Series1 refurbished
Brandname2 Productname2 Series2 black
```

We have three products that clearly represent one entity and forth that is
a different entity. Sample code to build clusters:

```go
package main

import (
	"fmt"
	"github.com/vharitonsky/glope"
)

func main() {
	transaction1 := glope.Transaction{
		//You can store anything here, an id or name to identify your transaction
		Instance: "Brandname1 Productname1 Series1 red",
		//Items is a series of unique tokens inside a transaction
		Items: []string{"Brandname1", "Productname1", "Series1", "red"},
	}
	transaction2 := glope.Transaction{
		Instance: "Brandname1 Productname1 Series1 blue",
		Items:    []string{"Brandname1", "Productname1", "Seris1", "blue"},
	}
	transaction3 := glope.Transaction{
		Instance: "Brandname1 Productname1 Series1 refurbished",
		Items:    []string{"Brandname1", "Productname1", "Series1", "refurbished"},
	}
	transaction4 := glope.Transaction{
		Instance: "Brandname2 Productname2 Series2 black",
		Items:    []string{"Brandname2", "Productname2", "Series2", "black"},
	}
	transactions := []*glope.Transaction{
		&transaction1,
		&transaction2,
		&transaction3,
		&transaction4,
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

```