package main

import "fmt"

func main() {
	qt1 := &Node{
		Val: true, IsLeaf: false,
		TopLeft:     &Node{Val: true, IsLeaf: true},
		TopRight:    &Node{Val: true, IsLeaf: true},
		BottomLeft:  &Node{Val: false, IsLeaf: true},
		BottomRight: &Node{Val: false, IsLeaf: true},
	}
	qt2 := &Node{
		Val: true, IsLeaf: false,
		TopLeft: &Node{Val: true, IsLeaf: true},
		TopRight: &Node{
			Val: true, IsLeaf: false,
			TopLeft:     &Node{Val: false, IsLeaf: true},
			TopRight:    &Node{Val: false, IsLeaf: true},
			BottomLeft:  &Node{Val: true, IsLeaf: true},
			BottomRight: &Node{Val: true, IsLeaf: true},
		},
		BottomLeft:  &Node{Val: true, IsLeaf: true},
		BottomRight: &Node{Val: false, IsLeaf: true},
	}
	fmt.Println(intersect(qt1, qt2))
}

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func (n *Node) String() string {
	l, v := 0, 0
	if n.IsLeaf {
		l = 1
	}
	if n.Val {
		v = 1
	}
	if !n.IsLeaf {
		return fmt.Sprintf("[%d %d] %s %s %s %s", l, v, n.TopLeft, n.TopRight, n.BottomLeft, n.BottomRight)
	} else {
		return fmt.Sprintf("[%d %d]", l, v)
	}
}

/**
 * Definition for a QuadTree node.
 * type Node struct {
 *     Val bool
 *     IsLeaf bool
 *     TopLeft *Node
 *     TopRight *Node
 *     BottomLeft *Node
 *     BottomRight *Node
 * }
 */

func intersect(quadTree1 *Node, quadTree2 *Node) *Node {
	if quadTree1 == nil {
		return quadTree2
	}
	if quadTree2 == nil {
		return quadTree1
	}
	if quadTree1.IsLeaf {
		if quadTree1.Val {
			return quadTree1
		}
		return quadTree2
	}
	if quadTree2.IsLeaf {
		if quadTree2.Val {
			return quadTree2
		}
		return quadTree1
	}
	tl := intersect(quadTree1.TopLeft, quadTree2.TopLeft)
	tr := intersect(quadTree1.TopRight, quadTree2.TopRight)
	bl := intersect(quadTree1.BottomLeft, quadTree2.BottomLeft)
	br := intersect(quadTree1.BottomRight, quadTree2.BottomRight)
	if tl.IsLeaf && tr.IsLeaf && bl.IsLeaf && br.IsLeaf && ((tl.Val && tr.Val && bl.Val && br.Val) || !(tl.Val || tr.Val || bl.Val || br.Val)) {
		return &Node{tl.Val && tr.Val && bl.Val && br.Val, true, nil, nil, nil, nil}
	}
	return &Node{tl.Val && tr.Val && bl.Val && br.Val, false, tl, tr, bl, br}
}
