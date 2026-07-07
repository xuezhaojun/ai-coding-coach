package lca_bst

import "testing"

func TestLowestCommonAncestor(t *testing.T) {
	tests := []struct {
		name   string
		vals   []int
		pVal   int
		qVal   int
		wantV  int
	}{
		{"example 1", []int{6, 2, 8, 0, 4, 7, 9, -101, -101, 3, 5}, 2, 8, 6},
		{"example 2", []int{6, 2, 8, 0, 4, 7, 9, -101, -101, 3, 5}, 2, 4, 2},
		{"root is lca", []int{2, 1, 3}, 1, 3, 2},
		{"same node", []int{2, 1, 3}, 1, 1, 1},
		{"right subtree", []int{6, 2, 8, 0, 4, 7, 9}, 7, 9, 8},
		{"deep nodes", []int{6, 2, 8, 0, 4, 7, 9, -101, -101, 3, 5}, 3, 5, 4},
	}

	// findNode finds a node by value in the tree.
	var findNode func(root *TreeNode, val int) *TreeNode
	findNode = func(root *TreeNode, val int) *TreeNode {
		if root == nil || root.Val == val {
			return root
		}
		if l := findNode(root.Left, val); l != nil {
			return l
		}
		return findNode(root.Right, val)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTreeHelper(tt.vals)
			p := findNode(root, tt.pVal)
			q := findNode(root, tt.qVal)
			got := lowestCommonAncestor(root, p, q)
			if got == nil || got.Val != tt.wantV {
				gotV := -1
				if got != nil {
					gotV = got.Val
				}
				t.Errorf("lowestCommonAncestor() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}
