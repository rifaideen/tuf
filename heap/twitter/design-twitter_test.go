package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {
	twitter := Constructor()
	twitter.PostTweet(1, 5)
	assert.Equal(t, []int{5}, twitter.GetNewsFeed(1))
	twitter.Follow(1, 2)
	twitter.PostTweet(2, 6)
	assert.Equal(t, []int{6, 5}, twitter.GetNewsFeed(1))
	twitter.Unfollow(1, 2)
	assert.Equal(t, []int{5}, twitter.GetNewsFeed(1))
}
