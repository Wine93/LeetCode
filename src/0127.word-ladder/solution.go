package solution

type Queue struct {
	prev *Queue
	next *Queue
}

func (q *Queue) Init() {
	q.prev = q
	q.next = q
}

func (h *Queue) Empty() bool {
	return h == h.prev
}

func ladderLength(beginWord string, endWord string, wordList []string) int {

}

/*

示例 1:

输入:
beginWord = "hit",
endWord = "cog",
wordList = ["hot","dot","dog","lot","log","cog"]

输出: 5

解释: 一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog",
     返回它的长度 5。
示例 2:

输入:
beginWord = "hit"
endWord = "cog"
wordList = ["hot","dot","dog","lot","log"]

输出: 0

解释: endWord "cog" 不在字典中，所以无法进行转换。
*/
