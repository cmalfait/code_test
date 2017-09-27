package main

import (
	"testing"
)

func TestGraph(t *testing.T) {
	d := New()
	d.AddNode("s1", 1)
	d.AddNode("s2", 2)
	d.AddNode("t1", 12)
	d.AddNode("t2", 13)
	d.AddEdge("s1", "t1")
	d.AddEdge("s1", "t2")
	d.AddEdge("s2", "t2")
}
