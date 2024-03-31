package alpha

import (
	"fmt"
	"strings"
)

// router 路由森林
type router struct {
	trees map[string]*node
}

func newRouter() *router {
	return &router{trees: make(map[string]*node)}
}

// addRouter 注册路由
func (r *router) addRouter(method string, pattern string, handleFunc HandleFunc) {
	fmt.Printf("add router %s - %s\n", method, pattern)
	if pattern == "" {
		panic("web: 路由不能为空")
	}

	// 获取根节点
	root, ok := r.trees[method]
	if !ok {
		// create root node
		// put root node into trees
		root = &node{
			part: "/",
		}
		r.trees[method] = root
	}
	// TODO 根路由 /
	if pattern == "/" {
		root.handleFunc = handleFunc
		return
	}

	if !strings.HasPrefix(pattern, "/") {
		panic("web: 路由必须以 / 开头")
	}
	if strings.HasSuffix(pattern, "/") {
		panic("web: 路由不准以 / 结尾")
	}

	// 切割 pattern
	parts := strings.Split(pattern[1:], "/")
	for _, part := range parts {
		if part == "" {
			panic("web: 路由不能来连续出现 /")
		}
		root = root.addNode(part)
	}
	root.handleFunc = handleFunc
}

func (r *router) getRouter(method string, pattern string) (*node, map[string]string, bool) {
	params := make(map[string]string)

	if pattern == "" {
		return nil, params, false
	}

	// 获取根节点
	root, ok := r.trees[method]
	if !ok {
		return nil, params, false
	}
	// TODO / 根路由
	if pattern == "/" {
		return root, params, true
	}

	// 切割 pattern
	parts := strings.Split(strings.Trim(pattern, "/"), "/")
	for _, part := range parts {
		if part == "" {
			return nil, params, false
		}
		root = root.getNode(part)
		if root == nil {
			return nil, params, false
		}
		if strings.HasPrefix(root.part, ":") {
			params[root.part[1:]] = part
		}
	}
	return root, params, root.handleFunc != nil
}

type node struct {
	part string
	// children 静态路由
	children map[string]*node
	// handleFunc 节点视图函数
	handleFunc HandleFunc
	// paramChild 参数路由
	paramChild *node
}

func (n *node) addNode(part string) *node {
	if strings.HasPrefix(part, ":") && n.paramChild == nil {
		n.paramChild = &node{part: part}
		return n.paramChild
	}

	if n.children == nil {
		n.children = make(map[string]*node)
	}
	child, ok := n.children[part]
	if !ok {
		child = &node{
			part: part,
		}
		n.children[part] = child
	}
	return child
}

func (n *node) getNode(part string) *node {
	if n.paramChild != nil {
		return n.paramChild
	}

	if n.children == nil {
		return nil
	}

	child, ok := n.children[part]
	if !ok {
		if n.paramChild != nil {
			return n.paramChild
		}
		return nil
	}
	return child
}
