package main

import "fmt"

// Node : A collection of node is a binary tree with integer values.
type Node struct {
	Left  *Node
	Value int32
	Right *Node
}

func main() {
	arr := [][]int32{
		{2, 3},
		{4, -1},
		{5, -1},
		{6, -1},
		{7, 8},
		{-1, 9},
		{-1, -1},
		{10, 11},
		{-1, -1},
		{-1, -1},
		{-1, -1},
	}

	q := []int32{2, 4}

	res := swapNodes(arr, q)

	for _, row := range res {
		for _, col := range row {
			fmt.Printf("%v ", col)
		}
		fmt.Println()
	}
}

func swapNodes(indexes [][]int32, queries []int32) [][]int32 {
	// Slice for containing the final result
	resultArr := [][]int32{}

	// Initializing the root of the node
	root := new(Node)

	// Build the tree
	tree := InsertLevelOrder(indexes, root, 0)

	// Looping as many as queries value
	for _, val := range queries {

		// Run the swap child function
		swappedTree := Swap(tree, val, 0)

		// Convert the swapped child node into array
		result := TreeToArrayInOrder(swappedTree, new([]int32))

		// Fill the final result slice with the array converted tree
		resultArr = append(resultArr, *result)
	}

	return resultArr
}

// InsertLevelOrder : Function for constructing the tree in level order
func InsertLevelOrder(arr [][]int32, node *Node, index int) *Node {

	// Check if arr slice is not empty
	if len(arr) == 0 || node == nil {

		// If empty return nil value
		return nil
	}

	// arrInteger is an array converted tree with the root value of 1
	arrInteger := []int32{1}

	// Populate the arrInteger
	for _, val := range arr {
		for _, el := range val {
			arrInteger = append(arrInteger, el)
		}
	}

	// Create the queue for node
	// Queue is a First In First Out
	treeNodeQueue := []*Node{}

	// Create queue for node value
	integerQueue := arrInteger[1:]

	// For the root node
	// Its value is always on arrInteger[0]
	node.Value = arrInteger[0]

	// Insert the root node to the node queue
	treeNodeQueue = append(treeNodeQueue, node)

	// Loop this logic as long as minimum there is a node
	// waiting in the node queue
	for len(treeNodeQueue) > 0 {

		// Initializing value for left and right child
		// -1 is a replacement value for nil
		leftVal := int32(-1)
		rightVal := int32(-1)

		// If the node value queue have a queue
		// Then assign the first value to the value for left node
		if len(integerQueue) > 0 {
			leftVal = integerQueue[0]

			// After assigning the first value to the left value
			// for the left node, pops that value out of the queue
			integerQueue = integerQueue[1:]
		}

		// After assign the left value for the left node
		// If it still have a queue waiting to assigned
		// Assign the first value to the right value
		// For the right node
		if len(integerQueue) > 0 {
			rightVal = integerQueue[0]

			// After assigning the first value to the right value
			// for the right node, pops that value out of the queue
			integerQueue = integerQueue[1:]
		}

		// Assign the first node in the queue
		currentNode := treeNodeQueue[0]

		// After assign the first node
		// Pops that node out of the queue
		treeNodeQueue = treeNodeQueue[1:]

		// If leftVal is not nil
		// -1 is a replacement value for nil
		if leftVal != -1 {

			// Create new node for the left node
			nodeLeft := new(Node)

			// Assign its value
			nodeLeft.Value = leftVal

			// Assign the left node to current node
			currentNode.Left = nodeLeft

			// And append the newly created node
			// Left node to the node queue for assigning its child
			treeNodeQueue = append(treeNodeQueue, nodeLeft)
		}

		// If rightVal is not nil
		// -1 is a replacement value for nil
		if rightVal != -1 {

			// Create new node for the left node
			nodeRight := new(Node)

			// Assign its value
			nodeRight.Value = rightVal

			// Assign the right node to current node
			currentNode.Right = nodeRight

			// And append the newly created node
			// Right node to the node queue for assigning its child
			treeNodeQueue = append(treeNodeQueue, nodeRight)
		}
	}

	return node
}

// TreeToArrayInOrder : Function for converting the tree to array using inorder traversal
// LEFT, ROOT, RIGHT
func TreeToArrayInOrder(tree *Node, result *[]int32) *[]int32 {
	// Check if the tree is not empty
	if tree != nil {
		// In Order Traversal
		// Left, Root, Right

		// Traversing the left child
		// until there are no left child anymore
		TreeToArrayInOrder(tree.Left, result)

		// Append the node value to the result slice
		*result = append(*result, tree.Value)

		// After traversing the left child and root
		// traverse the right child
		TreeToArrayInOrder(tree.Right, result)
	}

	return result
}

// Swap : Function for traverse the tree and swap its child nodes
func Swap(nodes *Node, query int32, depth int32) *Node {
	// The depth of the root is 1
	depth++

	// Check if there is a node
	if nodes != nil {

		// If the depth is equal to query
		// or equal to its multiple then swap its child nodes
		if depth%query == 0 {
			nodes.Left, nodes.Right = nodes.Right, nodes.Left
		}

		// Traversing using in order traversal
		// Left, Root, Right
		// Check if node have left child
		// If yes, traverse to that child node
		if nodes.Left != nil {
			Swap(nodes.Left, query, depth)
		}

		// Check if node have right child
		// If yes, traverse to that child node
		if nodes.Right != nil {
			Swap(nodes.Right, query, depth)
		}
	}

	return nodes
}
