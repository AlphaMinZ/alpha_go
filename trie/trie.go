package trie

import (
	"errors"
	"strings"
)

type Router struct {
	root map[string]*Node
}

// AddRouter 最开始的数据字符串
func (r *Router) AddRouter(pattern string, data string) {
	if r.root == nil {
		r.root = make(map[string]*Node)
	}

	root, ok := r.root["/"]
	if !ok {
		root = &Node{
			part: "/",
		}
		r.root["/"] = root
	}

	parts := strings.Split(strings.Trim(pattern, "/"), "/")
	for _, part := range parts {
		if part == "" {
			panic("pattern 不符合")
		}
		root = root.addNode(part)
	}

	root.data = data
}

func (r *Router) GetRouter(pattern string) (*Node, error) {
	root, ok := r.root["/"]
	// 创建根路由
	if !ok {
		return nil, errors.New("根节点不存在")
	}

	parts := strings.Split(strings.Trim(pattern, "/"), "/")
	for _, part := range parts {
		if part == "" {
			return nil, errors.New("pattern 格式不正确")
		}
		root = root.getNode(part)
		if root == nil {
			return nil, errors.New("pattern 不存在")
		}
	}

	return root, nil
}

type Node struct {
	// part 当前节点的唯一标识
	part string

	// children 维护子节点数据
	children map[string]*Node

	// data 当前节点需要保存的数据
	data string
}

func (n *Node) addNode(part string) *Node {
	if n.children == nil {
		n.children = make(map[string]*Node)
	}
	child, ok := n.children[part]
	if !ok {
		child = &Node{
			part: part,
		}
		n.children[part] = child
	}
	return child
}

func (n *Node) getNode(part string) *Node {
	if n.children == nil {
		return nil
	}

	child, ok := n.children[part]
	if !ok {
		return nil
	}
	return child
}
