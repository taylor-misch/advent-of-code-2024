package utilities

// Node represents a node in the AVL tree
type Node struct {
	Key    int
	Height int
	Left   *Node
	Right  *Node
}

// AVLTree represents the AVL tree
type AVLTree struct {
	Root *Node
}

// NewNode creates a new AVL tree node
func NewNode(key int) *Node {
	return &Node{Key: key, Height: 1}
}

// Insert inserts a key into the AVL tree and balances the tree
func (t *AVLTree) Insert(key int) {
	t.Root = insert(t.Root, key)
}

func insert(node *Node, key int) *Node {
	if node == nil {
		return NewNode(key)
	}

	if key < node.Key {
		node.Left = insert(node.Left, key)
	} else if key > node.Key {
		node.Right = insert(node.Right, key)
	} else {
		return node // Duplicate keys are not allowed
	}

	node.Height = 1 + max(height(node.Left), height(node.Right))

	balance := getBalance(node)

	// Left Left Case
	if balance > 1 && key < node.Left.Key {
		return rightRotate(node)
	}

	// Right Right Case
	if balance < -1 && key > node.Right.Key {
		return leftRotate(node)
	}

	// Left Right Case
	if balance > 1 && key > node.Left.Key {
		node.Left = leftRotate(node.Left)
		return rightRotate(node)
	}

	// Right Left Case
	if balance < -1 && key < node.Right.Key {
		node.Right = rightRotate(node.Right)
		return leftRotate(node)
	}

	return node
}

func height(node *Node) int {
	if node == nil {
		return 0
	}
	return node.Height
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getBalance(node *Node) int {
	if node == nil {
		return 0
	}
	return height(node.Left) - height(node.Right)
}

func rightRotate(y *Node) *Node {
	x := y.Left
	T2 := x.Right

	x.Right = y
	y.Left = T2

	y.Height = max(height(y.Left), height(y.Right)) + 1
	x.Height = max(height(x.Left), height(x.Right)) + 1

	return x
}

func leftRotate(x *Node) *Node {
	y := x.Right
	T2 := y.Left

	y.Left = x
	x.Right = T2

	x.Height = max(height(x.Left), height(x.Right)) + 1
	y.Height = max(height(y.Left), height(y.Right)) + 1

	return y
}
