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

func processLine (line string) {
	fmt.Printf("%s ", line)
	words := strings.Fields(line)
	fmt.Printf(" :length -> %d \n", len(words))
}

func main() {
	if len(os.Args) < 1 {
		fmt.Println("Invalid Arugument count.")
		fmt.Println("Please pass the file name to read.")
		os.Exit(128)
	}

	lines := readInput(os.Args[1])

	for i:=0; i < len(lines); i++ {
		processLine(lines[i])
	}
}
