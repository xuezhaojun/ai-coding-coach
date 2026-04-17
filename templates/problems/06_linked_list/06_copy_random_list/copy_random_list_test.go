package copy_random_list

import "testing"

// buildRandomList builds a RandomNode list from vals and randomIndices.
// randomIndices[i] is the index that node i's Random points to, or -1 for nil.
func buildRandomList(vals []int, randomIndices []int) *RandomNode {
	if len(vals) == 0 {
		return nil
	}
	nodes := make([]*RandomNode, len(vals))
	for i, v := range vals {
		nodes[i] = &RandomNode{Val: v}
	}
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Next = nodes[i+1]
	}
	for i, ri := range randomIndices {
		if ri >= 0 {
			nodes[i].Random = nodes[ri]
		}
	}
	return nodes[0]
}

func randomListToSlice(head *RandomNode) ([]int, []int) {
	// map node pointer to index
	indexMap := map[*RandomNode]int{}
	idx := 0
	for cur := head; cur != nil; cur = cur.Next {
		indexMap[cur] = idx
		idx++
	}
	var vals []int
	var randoms []int
	for cur := head; cur != nil; cur = cur.Next {
		vals = append(vals, cur.Val)
		if cur.Random != nil {
			randoms = append(randoms, indexMap[cur.Random])
		} else {
			randoms = append(randoms, -1)
		}
	}
	return vals, randoms
}

func TestCopyRandomList(t *testing.T) {
	tests := []struct {
		name          string
		vals          []int
		randomIndices []int
	}{
		{"nil list", nil, nil},
		{"single no random", []int{1}, []int{-1}},
		{"single self random", []int{1}, []int{0}},
		{"two nodes", []int{1, 2}, []int{1, 0}},
		{"three nodes mixed", []int{7, 13, 11}, []int{-1, 0, 2}},
		{"all random nil", []int{1, 2, 3, 4}, []int{-1, -1, -1, -1}},
		{"chain random", []int{1, 2, 3}, []int{2, 0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := buildRandomList(tt.vals, tt.randomIndices)
			copied := copyRandomList(head)

			if head == nil {
				if copied != nil {
					t.Errorf("expected nil, got non-nil")
				}
				return
			}

			gotVals, gotRandoms := randomListToSlice(copied)

			// verify values
			for i, v := range gotVals {
				if v != tt.vals[i] {
					t.Errorf("val[%d] = %d, want %d", i, v, tt.vals[i])
				}
			}
			// verify random indices
			for i, r := range gotRandoms {
				if r != tt.randomIndices[i] {
					t.Errorf("random[%d] = %d, want %d", i, r, tt.randomIndices[i])
				}
			}
			// verify deep copy (no shared nodes)
			origCur := head
			copyCur := copied
			for origCur != nil {
				if origCur == copyCur {
					t.Errorf("copy shares node with original at val=%d", origCur.Val)
				}
				origCur = origCur.Next
				copyCur = copyCur.Next
			}
		})
	}
}
