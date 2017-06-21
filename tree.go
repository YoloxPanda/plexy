package plexy

import (
	"fmt"
	"strings"
)

type node struct {
	parent   *node
	children []*node
	val      string
	handler  PlexyHandler
}

func (n *node) isPlaceholder() bool {
	// A placeholder begins with ':'
	return strings.HasPrefix(n.val, ":")
}

func (n *node) addChild(val string, handler PlexyHandler) {
	child := constructNode(val, handler)
	child.parent = n
	n.children = append(n.children, child)
}

func constructNode(s string, handler PlexyHandler) *node {
	split := cleanse(strings.Split(s, "/"))

	n := &node{
		parent:   nil,
		children: []*node{},
	}

	if len(split) == 1 {
		n.val = split[0]
		return n
	}

	cn := n
	for i, x := range split {
		cn.val = x

		if i != len(split)-1 {
			childNode := &node{
				parent:   cn,
				children: []*node{},
			}
			cn.children = append(cn.children, childNode)
			cn = childNode
		}
	}

	cn.handler = handler
	fmt.Println(cn.val)
	return n
}

func cleanse(s []string) []string {
	news := []string{}
	for _, x := range s {
		if strings.Replace(x, " ", "", -1) == "" {
			continue
		}
		news = append(news, x)
	}
	return news
}

type pathHandler struct {
	paths []*node
}

func (p *pathHandler) addPath(n *node) {
	p.paths = append(p.paths, n)
}

func (p *pathHandler) addPath2(val string, handler PlexyHandler) {
	// try to find a pre-existing node
	fn := p.findNode(val)

	if fn == nil {

	}
}

func (p *pathHandler) findNode(s string) *node {
	split := cleanse(strings.Split(s, "/"))
	depth := 0
	ca := p.paths
	var foundNode *node

	for i := 0; i < len(split); i++ {
		for _, x := range ca {
			if x.val == split[depth] {
				depth++
				ca = x.children
				foundNode = x
			}
		}
	}

	return foundNode
}

func (p *pathHandler) matchPath(path string) (*node, *Params) {
	split := cleanse(strings.Split(path, "/"))
	depth := 0
	// search for the top paths (all the parents)
	ca := p.paths

	params := newParams()
	var foundNode *node

	for i := 0; i < len(split); i++ {
		for _, x := range ca {
			// if it equals the path name OR is a placeholder
			if x.val == split[depth] || x.isPlaceholder() {

				if x.isPlaceholder() {
					params.params[x.val[1:]] = split[depth]
				}

				depth++
				ca = x.children
				foundNode = x
			}
		}
	}

	if depth != len(split) {
		fmt.Println("Depth:", depth, "Length:", len(split))
		return nil, nil
	}

	return foundNode, params
}

// func main() {
// 	ph := pathHandler{}
// 	n := constructNode("/user/:username")
// 	n2 := constructNode("/users")
// 	n3 := constructNode("/:search")

// 	ph.addPath(n)
// 	ph.addPath(n2)
// 	ph.addPath(n3)

// 	fn := ph.pathMatch("/user/:username/test")
// 	fmt.Println(fn.val)
// }
