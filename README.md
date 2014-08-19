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

import(
	"fmt"
	"github.com/vhartionsky/glope"
)

func main(){
	transaction1 := glope.Transaction{
		Instance:"Brandname1 Productname1 Series1 red" //You can store anything here, an id or name to identify your transaction
		Items:[]string{"Brandname1", "Productname1", "Series1", "red"}
	}
	transaction2 := glope.Transaction{
		Instance:"Brandname1 Productname1 Series1 blue"
		Items:[]string{"Brandname1", "Productname1", "Series1", "blue"}
	}
	transaction3 := glope.Transaction{
		Instance:"Brandname1 Productname1 Series1 refurbished"
		Items:[]string{"Brandname1", "Productname1", "Series1", "refurbished"}
	}
	transaction4 := glope.Transaction{
		Instance:"Brandname2 Productname2 Series2 black" 
		Items:[]string{"Brandname2", "Productname2", "Series2", "black"}
	}
	transactions = []*Transaction{&transaction1, &transaction2, &transaction3, &transaction4}
	//Clusterize takes two parameters: list of pointers to transactions and repulsion.
	//Repulsion parameter is a threshold of likeness for transactions, the higher, the more precise.
	//In real life, repulsion depends on the amount of items in your transaction, the more items you have, the higher the repulsion should be.
	//But its up to you to conclude by trial and error the value of repulsion parameter. Default is 4.

	clusters := glope.Clusterize(transactions, 4)
	for _, cluster := range clusters{
		for _, transaction := range cluster.Transactions{
			fmt.Printf("%s in %v", transaction.Instance, cluster)
		}
	}
}
```