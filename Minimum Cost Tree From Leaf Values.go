type Node struct {
	value         int
	maxChildValue int
	leaf          bool
}

func max(i, j Node) int {
	if i.maxChildValue > j.maxChildValue {
		return i.maxChildValue
	}
	return j.maxChildValue
}

// modified huffman tree
// not corrent, I don't understand its requirement, it need me remain the array order
func mctFromLeafValues(arr []int) int {
	tree := []Node{}
	for i := 0; i < len(arr); i++ {
		tree = append(tree, Node{value: arr[i], maxChildValue: arr[i], leaf: true})
	}
	sort.Slice(tree, func(i, j int) bool {
		return tree[i].maxChildValue < tree[j].maxChildValue
	})

	for i := 0; i < len(arr)-1; i++ {
		tree = append(tree, Node{leaf: false, value: tree[0].maxChildValue * tree[1].maxChildValue,
			maxChildValue: max(tree[0], tree[1])})
		sort.Slice(tree, func(i, j int) bool {
			return tree[i].maxChildValue < tree[j].maxChildValue
		})
	}
	result := 0
	for i := 0; i < 2*len(tree)-1; i++ {
		if !tree[i].leaf {
			result += tree[i].value
		}
	}
	return result
}

// Greedy algorithm, O(N^2)
func mctFromLeafValues(arr []int) int {
	result := 0
	tree := make([]int, len(arr))
	copy(tree, arr)
	for len(tree) > 1 {
		min := 2<<30 - 1
		var l int
		for i := 1; i < len(tree); i++ {
			if tree[i-1]*tree[i] < min {
				min = tree[i-1] * tree[i]
				if tree[i-1] < tree[i] {
					l = i - 1
				} else {
					l = i
				}
			}
		}
		result += min
		tree = append(tree[:l], tree[l+1:]...)
	}
	return result
}