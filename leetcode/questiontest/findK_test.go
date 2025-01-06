package questiontest

import (
	"fmt"
	"github.com/hutaochu/note/leetcode/questionbank"
	"math/rand"
	"testing"
)

func TestBinaryFindTarget(t *testing.T) {
	nums := []int{
		0, 0, 0, 0, 0,
		1, 2, 3, 4, 5,
		7, 7, 7, 7, 8,
		8, 8, 9, 9, 10,
	}
	nums = []int{0, 1, 2}
	target := 0
	index := questionbank.BinaryFindTarget(nums, target)
	fmt.Println("index: ", index)
	fmt.Println(nums[index])
}

func TestMergeListNode(t *testing.T) {
	l1 := &questionbank.ListNode{
		Var: 1,
	}
	l1.Next = &questionbank.ListNode{
		Var: 4,
	}
	l1.Next.Next = &questionbank.ListNode{
		Var: 9,
	}
	l2 := &questionbank.ListNode{
		Var: 3,
	}
	l2.Next = &questionbank.ListNode{
		Var: 6,
	}
	l2.Next.Next = &questionbank.ListNode{
		Var: 7,
	}
	l2.Next.Next.Next = &questionbank.ListNode{
		Var: 9,
	}
	fmt.Println("l1: ")
	l1.Print()
	fmt.Println("\nl2: ")
	l2.Print()
	l3 := questionbank.MergeList(l1, l2)
	fmt.Println("\nl3: ")
	l3.Print()
}

func TestPartitionListNode(t *testing.T) {
	l1 := &questionbank.ListNode{
		Var: 1,
	}
	p := l1
	for i := 0; i < 10; i++ {
		l := &questionbank.ListNode{
			Var: rand.Intn(20),
		}
		p.Next = l
		p = p.Next
	}
	fmt.Println("l1: ")
	l1.Print()
	l2 := questionbank.PartitionList(l1, 5)
	fmt.Println("\nl2: ")
	l2.Print()
}

func TestMergeKList(t *testing.T) {
	l1 := &questionbank.ListNode{
		Var: 1,
		Next: &questionbank.ListNode{
			Var: 3,
			Next: &questionbank.ListNode{
				Var: 7,
				Next: &questionbank.ListNode{
					Var: 8,
				},
			},
		},
	}
	l2 := &questionbank.ListNode{
		Var: 3,
		Next: &questionbank.ListNode{
			Var: 9,
			Next: &questionbank.ListNode{
				Var: 17,
				Next: &questionbank.ListNode{
					Var: 18,
					Next: &questionbank.ListNode{
						Var: 18,
					},
				},
			},
		},
	}
	l3 := &questionbank.ListNode{
		Var: 1,
		Next: &questionbank.ListNode{
			Var: 5,
			Next: &questionbank.ListNode{
				Var: 27,
				Next: &questionbank.ListNode{
					Var: 38,
				},
			},
		},
	}
	fmt.Println("l1: ")
	l1.Print()
	fmt.Println("\nl2: ")
	l2.Print()
	fmt.Println("\nl3: ")
	l3.Print()
	r := questionbank.MergeKList([]*questionbank.ListNode{l1, l2, l3})
	fmt.Println("\nresult: ")
	r.Print()
}

func TestFindKthLastInList(t *testing.T) {
	l1 := &questionbank.ListNode{
		Var: 3,
		Next: &questionbank.ListNode{
			Var: 9,
			Next: &questionbank.ListNode{
				Var: 17,
				Next: &questionbank.ListNode{
					Var: 18,
					Next: &questionbank.ListNode{
						Var: 20,
					},
				},
			},
		},
	}
	l := questionbank.FindKthLastInList(l1, 6)
	l.Print()

}

func TestFindMidNode(t *testing.T) {
	l1 := &questionbank.ListNode{
		Var: 3,
		Next: &questionbank.ListNode{
			Var: 9,
			Next: &questionbank.ListNode{
				Var: 17,
				Next: &questionbank.ListNode{
					Var: 18,
					Next: &questionbank.ListNode{
						Var: 20,
						Next: &questionbank.ListNode{
							Var: 21,
						},
					},
				},
			},
		},
	}
	l := questionbank.FindMidNode(l1)
	l.Print()
}

func TestHasCycle(t *testing.T) {
	l1 := &questionbank.ListNode{
		Var: 3,
		Next: &questionbank.ListNode{
			Var: 9,
			Next: &questionbank.ListNode{
				Var: 17,
				Next: &questionbank.ListNode{
					Var: 18,
					Next: &questionbank.ListNode{
						Var: 20,
						Next: &questionbank.ListNode{
							Var: 21,
						},
					},
				},
			},
		},
	}
	hasCycle := questionbank.HasCycle(l1)
	fmt.Println("hasCycle: ", hasCycle)

	l2 := &questionbank.ListNode{
		Var: 3,
	}
	n1 := &questionbank.ListNode{
		Var: 9,
	}
	n2 := &questionbank.ListNode{
		Var: 17,
	}
	n3 := &questionbank.ListNode{
		Var: 18,
	}
	n4 := &questionbank.ListNode{
		Var: 20,
	}
	n5 := &questionbank.ListNode{
		Var: 21,
	}
	l2.Next = n1
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n5.Next = n3

	hasCycle = questionbank.HasCycle(l2)
	fmt.Println("hasCycle: ", hasCycle)

	cycleNode := questionbank.FindCycleStartNode(l2)
	if cycleNode != nil {
		fmt.Println("cycle start node: ", cycleNode.Var)
	}
}

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{1, 2, 3, 4, 4, 4, 5, 6, 6, 7}
	k := questionbank.RemoveDuplicates(nums)
	fmt.Println(nums[:k])
}

func TestFindKSum(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 3, 5, 6}
	fmt.Println(questionbank.FindKLargestSums(a, b, 3))
}
