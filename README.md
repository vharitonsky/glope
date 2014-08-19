# Glope - a Go implementation of Clope algorythm of clusterization.

What is this?
=========

A library that implements Clope algorythm of clusterization.
You may need it to clusterize an array of transactions with overlapping items.
Like, for example, names of products that differ by minor tokens.

Example:
======

```txt
Brandname1 Productname1 Series1 red
Brandname1 Productname1 Series1 blue
Brandname1 Productname1 Series1 refurbished
Brandname2 Productname2 Series2 black
```

We have three products that clearly represent one entity and forth that is
a different entity, here is a sample code:

```go
package main

import (
	"fmt"
	"github.com/vharitonsky/glope"
)

func main() {
	transaction1 := glope.Transaction{
		Instance: "Brandname1 Productname1 Series1 color2 32Gb", //You can store anything here, an id or name to identify your transaction
		Items:    []string{"Brandname1", "Productname1", "Series1", "color2", "32Gb"},
	}
	transaction2 := glope.Transaction{
		Instance: "Brandname1 Productname1 Series1 color2 32Gb",
		Items:    []string{"Brandname1", "Productname1", "Seris1", "color2", "16Gb"},
	}
	transaction3 := glope.Transaction{
		Instance: "Brandname1 Productname1 Series1 color1 12Gb",
		Items:    []string{"Brandname1", "Productname1", "Series1", "color2", "12Gb"},
	}
	transaction4 := glope.Transaction{
		Instance: "Brandname2 Productname2 Series2 black",
		Items:    []string{"Brandname2", "Productname2", "Series2", "black"},
	}
	transactions := []*glope.Transaction{&transaction1, &transaction2, &transaction3, &transaction4}
	//Clusterize takes two parameters: list of pointers to transactions and repulsion.
	//Repulsion parameter is a threshold of likeness for transactions, the higher, the more precise.
	//In real life, repulsion depends on the amount of items in your transaction, the more items you have, the higher the repulsion should be.
	//But its up to you to conclude by trial and error the value of repulsion parameter. Default is 4.

	clusters := glope.Clusterize(transactions, 2)
	for _, cluster := range clusters {
		for _, transaction := range cluster.Transactions {
			fmt.Printf("%s in %v\n", transaction.Instance, cluster)
		}
	}
	/*
		Brandname1 Productname1 Series1 color2 32Gb in [Cluster 0]
		Brandname1 Productname1 Series1 color2 32Gb in [Cluster 0]
		Brandname1 Productname1 Series1 color1 16Gb in [Cluster 0]
		Brandname2 Productname2 Series2 black in [Cluster 1]
	*/
}
```