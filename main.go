// Chad Malfait
// Oct 2, 2017
// Program to install / remove system packages along with their dependencies

package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strings"
)

//map of nodes.  this holds our Graph
type Graph struct {
	nodes map[string]*Node
}

//a vertix/node of our graph
type Node struct {
	name string
	value string
	children []*Node
}

//Create a new graph
func New() *Graph {
	this := new(Graph)
	this.nodes = make(map[string]*Node)
	return this
}

//Process input file
func readInput (path string) ([]string) {
	var lines []string
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

/*
func processLine (line string) {
	fmt.Printf("%s ", line)
	words := strings.Fields(line)
	fmt.Printf(" :length -> %d \n", len(words))
	fmt.Printf(" :one %s \n", words[0])
	fmt.Printf(" :one %s \n", words[1])
}
*/

func addEdge(mystring string) {
	fmt.Printf("AddEdge %s \n", mystring)
}

func addNode(mystring string) {
	fmt.Printf("AddNode %s \n", mystring)
}

func removeEdge(mystring string) {
	fmt.Printf("RemoveEdge %s \n", mystring)
}

func removeNode(mystring string) {
	fmt.Printf("RemoveNode %s \n", mystring)
}

func listNodes(mystring string) {
	fmt.Printf("list %s \n", mystring)
}

func main() {
	if len(os.Args) < 1 {
		fmt.Println("Invalid Arugument count.")
		fmt.Println("Please pass the file name to read.")
		os.Exit(128)
	}

	lines := readInput(os.Args[1])

	for i:=0; i < len(lines); i++ {
		words := strings.Fields(lines[i])

		switch words[0] {
        case "DEPEND":
          addEdge(lines[i])
        case "INSTALL":
          addNode(lines[i])
        case "REMOVE":
          removeEdge(lines[i])
          removeNode(lines[i])
        case "LIST":
          listNodes(lines[i])
        case "END":
          fmt.Println("END")
          os.Exit(128)
        default:
          fmt.Println("Invalid Arguemnt")
          fmt.Println("Valid input is: DEPEND, INSTALL, REMOVE, LIST, END")
        }
	}
}
