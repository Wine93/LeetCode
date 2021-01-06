package solution

import "sort"

type Pair struct {
	Key   int
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value > p[j].Value }

func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		count[nums[i]]++
	}

	p := make(PairList, len(count))
	for k, v := range count {
		p = append(p, Pair{k, v})
	}

	sort.Sort(p)

	o := []int{}
	for i := 0; i < k; i++ {
		o = append(o, p[i].Key)
	}

	return o
}
