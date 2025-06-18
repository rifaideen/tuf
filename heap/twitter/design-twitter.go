package main

type Tweet struct {
	Id       int
	UserId   int
	Priority int
}

type User struct {
	Id         int
	Followings map[int]*User // map[user-id] = user object
	Newsfeed   Newsfeed      // list of news feed
}

type Newsfeed struct {
	Tweets []Tweet
}

type Twitter struct {
	tweets  *maxheap
	users   map[int]*User // map[user-id] = user object
	counter int
}

func Constructor() Twitter {
	return Twitter{
		tweets: &maxheap{
			size:   0,
			tweets: []Tweet{},
		},
		users: make(map[int]*User),
	}
}

func (t *Twitter) PostTweet(userId int, tweetId int) {
	t.counter++
	t.tweets.push(Tweet{
		Id:       tweetId,
		UserId:   userId,
		Priority: t.counter,
	})

	if _, ok := t.users[userId]; !ok {
		t.users[userId] = &User{
			Id:         userId,
			Followings: make(map[int]*User),
		}
	}
}

func (t *Twitter) GetNewsFeed(userId int) []int {
	cycle := 10
	tweets := []Tweet{}
	ids := []int{}
	follower := t.users[userId]

	if follower == nil {
		return ids
	}

	followings := follower.Followings

	for cycle > 0 && t.tweets.size > 0 {
		tweet := t.tweets.Pop()
		// regardless of whether it's match or not, we need to put
		tweets = append(tweets, tweet)

		// check if the tweet author is followed by the current user or the tweet is by the author
		if tweet.UserId == userId || (followings[tweet.UserId] != nil) {
			ids = append(ids, tweet.Id)
			cycle--
		}
	}

	for _, tweet := range tweets {
		t.tweets.push(tweet)
	}

	return ids
}

func (t *Twitter) Follow(followerId int, followeeId int) {
	var follower *User
	var followee *User
	var ok bool

	if follower, ok = t.users[followerId]; !ok {
		follower = &User{
			Id:         followerId,
			Followings: make(map[int]*User),
		}

		t.users[followerId] = follower
	}

	if followee, ok = t.users[followeeId]; !ok {
		followee = &User{
			Id:         followeeId,
			Followings: make(map[int]*User),
		}

		t.users[followeeId] = followee
	}

	follower.Followings[followeeId] = followee
}

func (t *Twitter) Unfollow(followerId int, followeeId int) {
	if follower, ok := t.users[followerId]; ok {
		delete(follower.Followings, followeeId)
	}
}

type maxheap struct {
	tweets []Tweet
	size   int
}

func (h *maxheap) push(tweet Tweet) {
	h.tweets = append(h.tweets, tweet)
	h.Up(h.size)
	h.size++
}

func (h *maxheap) Pop() Tweet {
	if h.size == 0 {
		panic("no tweets found")
	}

	tweet := h.tweets[0]

	h.size--
	h.tweets[0] = h.tweets[h.size]
	h.tweets = h.tweets[:h.size]

	h.Down(0)

	return tweet
}

func (h *maxheap) Up(index int) {
	if index > 0 {
		parent := (index - 1) / 2

		if h.tweets[index].Priority > h.tweets[parent].Priority {
			h.swap(index, parent)
			h.Up(parent)
		}
	}
}

func (h *maxheap) Down(index int) {
	left := (2 * index) + 1
	right := left + 1
	largest := index

	if left < h.size && h.tweets[left].Priority > h.tweets[largest].Priority {
		largest = left
	}

	if right < h.size && h.tweets[right].Priority > h.tweets[largest].Priority {
		largest = right
	}

	if largest != index {
		h.swap(index, largest)
		h.Down(largest)
	}
}

func (h *maxheap) swap(i, j int) {
	h.tweets[i], h.tweets[j] = h.tweets[j], h.tweets[i]
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
