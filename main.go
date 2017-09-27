// Chad Malfait
// Sept 26, 2017
// Program to install system packages along with their dependencies

package main

import (
	"fmt"
	"os"
)

type Graph struct {
	nodes map[string]*Node
}

type Node struct {
	name string
	value string
	cntr int
	children []*Node
}

func main() {
	if len(os.Args) < 1 {
		fmt.Println("Invalid Arugument count.")
		fmt.Println("Valid input is DEPEND, INSTALL, REMOVE, LIST, END")
		os.Exit(128)
	}

	switch os.Args[1] {
	case "DEPEND":
		addEdge("DEPEND")
	case "INSTALL":
		addNode("INSTALL")
	case "REMOVE":
		removeEdge("REMOVE")
		removeNode("REMOVE")
	case "LIST":
		listNodes("LIST")
	case "END":
		fmt.Println("END")
		os.Exit(128)
	default:
		fmt.Println("Invalid Arguemnt")
		fmt.Println("Valid input is: DEPEND, INSTALL, REMOVE, LIST, END")
	}
}

//Create a new graph
func New() *Graph {
	this := new(Graph)
	this.nodes = make(map[string]*Node)
	return this
}

//Function to add a node to our graph
func (this *Graph) AddNode(name string, value string) *Node {
	node := &Node{name: name, value: value}
	this.nodes[name] = node
	return node
}

//Function to add an edge to our node
func (this *Graph) AddEdge(from, to string) {
	fromNode := this.nodes[from]
	toNode := this.nodes[to]
	fromNode.children = append(fromNode.children, toNode)
	toNode.cntr++
}

//Function to remove an edge to our graph
func removeEdge(junk string) {
	fmt.Printf("hello %s", junk)
}

//Function to remove a node to our graph
func removeNode(junk string) {
	fmt.Printf("hello %s", junk)
}

//Function to list a nodes in our graph
func listNodes(junk string) {
	fmt.Printf("hello %s", junk)
}

