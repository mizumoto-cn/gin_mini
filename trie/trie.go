package trie

type node struct {
	// route waiting to be matched
	Path string
	// Part of the route
	Part string
	// Children nodes
	Children map[string]*node
	// precise matching or not, for dynamic route matching
	IsWild bool
}

// matchChild returns the first child node that matches
func (n *node) MatchChild(part string) *node {
	for _, child := range n.Children {
		if child.Part == part || child.IsWild {
			return child
		}
	}
	return nil
}

// matchChildren returns all children nodes that matches
func (n *node) MatchChildren(part string) []*node {
	nodes := make(*[]*node, 0)
	for _, child := range n.Children {
		if child.Part == part || child.IsWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// Insert inserts a new node as a child node
