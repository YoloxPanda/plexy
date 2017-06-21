package plexy

import "strings"

type node struct {
	parent   *node
	children []*node
	val      string
	handler  PlexyHandler
}

func (n *node) addChild(val string) {
	child := constructNode(val)
	child.parent = n
	n.children = append(n.children, child)
}

func constructNode(s string) *node {
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
	for _, x := range split {
		cn.val = x
		childNode := &node{
			parent:   cn,
			children: []*node{},
		}
		cn.children = append(cn.children, childNode)
		cn = childNode
	}

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
