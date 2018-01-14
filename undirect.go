package main

import (
	"os"
	"os/exec"

	"github.com/awalterschulze/gographviz"
)

func main() {
	// draw the undirected graph.
	g := gographviz.NewGraph()
	if err := g.SetName("G"); err != nil {
		panic(err)
	}
	// if err := g.SetDir(true); err != nil {
	// 	panic(err)
	// }
	if err := g.AddAttr("G", "bgcolor", "\"#343434\""); err != nil {
		panic(err)
	}
	if err := g.AddAttr("G", "layout", "circo"); err != nil {
		panic(err)
	}

	// configuration for nodes
	nodeAttrs := make(map[string]string)
	nodeAttrs["colorscheme"] = "rdylgn11"
	nodeAttrs["style"] = "\"solid,filled\""
	nodeAttrs["fontcolor"] = "6"
	nodeAttrs["fontname"] = "\"Migu 1M\""
	nodeAttrs["color"] = "7"
	nodeAttrs["fillcolor"] = "11"
	nodeAttrs["shape"] = "doublecircle"

	// make the nodes.
	// nodes := make(map[string]bool)
	// for f := range flows {
	// 	if !nodes[f.Src().String()] {
	// 		nodes[f.Src().String()] = true
	// 	}
	// 	if !nodes[f.Dst().String()] {
	// 		nodes[f.Dst().String()] = true
	// 	}
	// }

	// add nodes to the graph "G"
	nodes := []string{"A", "B", "C", "D", "E", "F"}
	for _, node := range nodes {
		if err := g.AddNode("G", node, nodeAttrs); err != nil {
			panic(err)
		}
	}

	// configuration for edges
	edgeAttrs := make(map[string]string)
	edgeAttrs["color"] = "white"

	// connect the nodes
	edges := [][]string{
		{"A", "F"},
		{"A", "B"},
		{"B", "C"},
		{"C", "D"},
		{"D", "E"},
		{"E", "B"},
	}
	for _, edge := range edges {
		err := g.AddEdge(edge[0], edge[1], false, edgeAttrs)
		if err != nil {
			panic(err)
		}
	}

	// save the graph as dotfile.
	s := g.String()
	dotfile := "./undigraph.dot"
	pngfile := "./undigraph.png"
	file, err := os.Create(dotfile)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.Write([]byte(s))

	if err = exec.Command("dot", "-T", "png",
		dotfile, "-o", pngfile).Run(); err != nil {
		panic(err)
	}

	if err = exec.Command("open", pngfile).Run(); err != nil {
		panic(err)
	}

}
