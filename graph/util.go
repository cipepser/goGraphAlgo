package graph

import (
	"os"
	"os/exec"
	"strconv"

	"github.com/awalterschulze/gographviz"
)

// Visualize create `gv.png` in your directory, which is generated by the graph `g`.
func (g *graph) Visualize() error {
	gv := gographviz.NewGraph()
	if err := gv.SetName("G"); err != nil {
		return err
	}

	if err := gv.SetDir(g.isDirected); err != nil {
		return err
	}

	if err := gv.AddAttr("G", "bgcolor", "\"#343434\""); err != nil {
		return err
	}
	if err := gv.AddAttr("G", "layout", "circo"); err != nil {
		return err
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

	// add vertices as nodes
	for _, v := range g.GetVertices() {
		if err := gv.AddNode("G", strconv.FormatUint(uint64(v), 10), nodeAttrs); err != nil {
			return err
		}
	}

	// configuration for edges
	edgeAttrs := make(map[string]string)
	edgeAttrs["color"] = "white"

	// add nodes
	for _, e := range g.GetEdges() {
		if err := gv.AddEdge(
			strconv.FormatUint(uint64(e.From), 10),
			strconv.FormatUint(uint64(e.To), 10),
			g.isDirected, edgeAttrs); err != nil {
			return err
		}
	}

	// save the graph as dotfile.
	s := gv.String()
	dotfile := "./gv.dot"
	pngfile := "./gv.png"
	file, err := os.Create(dotfile)
	if err != nil {
		return err
	}
	defer file.Close()
	file.Write([]byte(s))

	// generate png image by graphviz
	if err = exec.Command("dot", "-T", "png",
		dotfile, "-o", pngfile).Run(); err != nil {
		return err
	}

	// automatically open the generated png file
	// if err = exec.Command("open", pngfile).Run(); err != nil {
	// 	return err
	// }

	return nil
}