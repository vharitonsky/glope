package main

import (
	"math"
)

type Cluster struct {
	n         float64
	w         float64
	s         float64
	instances map[string]bool
	occ       map[string]int
}

func (c *Cluster) addItem(item string) {
	val, found := c.occ[item]
	if !found {
		c.occ[item] = 1
	} else {
		c.occ[item] = val + 1
	}
}

func (c *Cluster) deltaAdd(items []string, r float64) float64 {
	sNew := c.s + float64(len(items))
	wNew := c.w
	for _, item := range items {
		if _, found := c.occ[item]; !found {
			wNew++
		}
	}
	if c.n == 0 {
		return sNew * (c.n + 1) / math.Pow(wNew, r)
	} else {
		profit := c.s * c.n / math.Pow(c.w, r)
		profitNew := sNew * (c.n + 1) / math.Pow(wNew, r)
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

func (c *Cluster) addInstance(instance string, items []string) {
	for _, item := range items {
		c.addItem(item)
	}
	c.w = float64(len(c.occ))
	c.n++
	c.instances[instance] = true
}

func (c *Cluster) removeInstance(instance string, items []string) {
	for _, item := range items {
		c.removeItem(item)
	}
	c.w = float64(len(c.occ))
	c.n++
	delete(c.instances, instance)
}

func main() {

}
