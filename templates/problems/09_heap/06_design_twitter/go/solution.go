package design_twitter

import "container/heap"

// SolveTwitter implements the Twitter design.
type SolveTwitter struct {
	time      int
	tweets    map[int][]solveTweet
	following map[int]map[int]bool
}

type solveTweet struct {
	time    int
	tweetId int
}

func SolveConstructorTwitter() SolveTwitter {
	return SolveTwitter{
		tweets:    make(map[int][]solveTweet),
		following: make(map[int]map[int]bool),
	}
}

// Time: O(1).
func (t *SolveTwitter) PostTweet(userId int, tweetId int) {
	t.time++
	t.tweets[userId] = append(t.tweets[userId], solveTweet{time: t.time, tweetId: tweetId})
}

// GetNewsFeed returns the 10 most recent tweet IDs.
// Time: O(k log f) where k=10, f=number of followees.
func (t *SolveTwitter) GetNewsFeed(userId int) []int {
	h := &solveTweetHeap{}
	// include self
	users := []int{userId}
	if followees, ok := t.following[userId]; ok {
		for fid := range followees {
			users = append(users, fid)
		}
	}
	for _, uid := range users {
		tweets := t.tweets[uid]
		if len(tweets) > 0 {
			idx := len(tweets) - 1
			heap.Push(h, solveFeedItem{tweet: tweets[idx], userId: uid, idx: idx})
		}
	}
	var result []int
	for h.Len() > 0 && len(result) < 10 {
		item := heap.Pop(h).(solveFeedItem)
		result = append(result, item.tweet.tweetId)
		if item.idx > 0 {
			next := item.idx - 1
			heap.Push(h, solveFeedItem{
				tweet:  t.tweets[item.userId][next],
				userId: item.userId,
				idx:    next,
			})
		}
	}
	return result
}

// Time: O(1).
func (t *SolveTwitter) Follow(followerId int, followeeId int) {
	if t.following[followerId] == nil {
		t.following[followerId] = make(map[int]bool)
	}
	t.following[followerId][followeeId] = true
}

// Time: O(1).
func (t *SolveTwitter) Unfollow(followerId int, followeeId int) {
	if t.following[followerId] != nil {
		delete(t.following[followerId], followeeId)
	}
}

type solveFeedItem struct {
	tweet  solveTweet
	userId int
	idx    int
}

type solveTweetHeap []solveFeedItem

func (h solveTweetHeap) Len() int            { return len(h) }
func (h solveTweetHeap) Less(i, j int) bool   { return h[i].tweet.time > h[j].tweet.time }
func (h solveTweetHeap) Swap(i, j int)        { h[i], h[j] = h[j], h[i] }
func (h *solveTweetHeap) Push(x interface{})  { *h = append(*h, x.(solveFeedItem)) }
func (h *solveTweetHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
