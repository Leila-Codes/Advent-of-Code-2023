package documents

import (
	"bufio"
	"os"
	"strings"
)

type WastelandMap struct {
	Instructions string
	Nodes        map[string]*MapNode
}

func (wm *WastelandMap) BuildLinks() {
	for _, node := range wm.Nodes {
		node.Left = wm.Nodes[node.leftName]
		node.Right = wm.Nodes[node.rightName]
	}
}

func (wm *WastelandMap) FindAll(test func(string) bool) (matches []*MapNode) {
	matches = make([]*MapNode, 0)

	for key := range wm.Nodes {
		if test(key) {
			matches = append(matches, wm.Nodes[key])
		}
	}

	return matches
}

type MapNode struct {
	Value               string
	leftName, rightName string
	Left, Right         *MapNode
}

func nodeFromString(line string, wasteMap *WastelandMap) {
	var (
		parts    = strings.Split(line, "=")
		nodeName = strings.TrimSpace(parts[0])

		links          = strings.TrimSpace(parts[1])
		linkSplitIndex = strings.Index(links, ", ")
	)

	wasteMap.Nodes[nodeName] = &MapNode{
		Value:     nodeName,
		leftName:  links[1:linkSplitIndex],
		rightName: links[linkSplitIndex+2 : len(links)-1],
	}
}

func MapFromFile(filePath string) *WastelandMap {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	wastelandMap := &WastelandMap{Nodes: make(map[string]*MapNode)}

	reader := bufio.NewScanner(f)
	reader.Scan()
	wastelandMap.Instructions = strings.TrimSpace(reader.Text())

	for reader.Scan() {
		if len(reader.Text()) == 0 {
			continue
		}
		nodeFromString(reader.Text(), wastelandMap)
	}

	wastelandMap.BuildLinks()

	return wastelandMap
}
