package design_twitter

import (
	"reflect"
	"testing"
)

func TestTwitter(t *testing.T) {
	t.Run("basic post and feed", func(t *testing.T) {
		tw := ConstructorTwitter()
		tw.PostTweet(1, 5)
		got := tw.GetNewsFeed(1)
		want := []int{5}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("GetNewsFeed(1) = %v, want %v", got, want)
		}
	})

	t.Run("follow and see followee tweets", func(t *testing.T) {
		tw := ConstructorTwitter()
		tw.PostTweet(1, 5)
		tw.PostTweet(2, 6)
		tw.Follow(1, 2)
		got := tw.GetNewsFeed(1)
		want := []int{6, 5}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("GetNewsFeed(1) = %v, want %v", got, want)
		}
	})

	t.Run("unfollow removes tweets from feed", func(t *testing.T) {
		tw := ConstructorTwitter()
		tw.PostTweet(1, 5)
		tw.PostTweet(2, 6)
		tw.Follow(1, 2)
		tw.Unfollow(1, 2)
		got := tw.GetNewsFeed(1)
		want := []int{5}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("GetNewsFeed(1) = %v, want %v", got, want)
		}
	})

	t.Run("feed limited to 10 most recent", func(t *testing.T) {
		tw := ConstructorTwitter()
		for i := 1; i <= 12; i++ {
			tw.PostTweet(1, i)
		}
		got := tw.GetNewsFeed(1)
		if len(got) != 10 {
			t.Errorf("GetNewsFeed returned %d tweets, want 10", len(got))
		}
		want := []int{12, 11, 10, 9, 8, 7, 6, 5, 4, 3}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("GetNewsFeed(1) = %v, want %v", got, want)
		}
	})

	t.Run("empty feed for new user", func(t *testing.T) {
		tw := ConstructorTwitter()
		got := tw.GetNewsFeed(1)
		if len(got) != 0 {
			t.Errorf("GetNewsFeed(1) = %v, want empty", got)
		}
	})

	t.Run("multiple followees merged feed", func(t *testing.T) {
		tw := ConstructorTwitter()
		tw.PostTweet(2, 10)
		tw.PostTweet(3, 20)
		tw.PostTweet(2, 30)
		tw.Follow(1, 2)
		tw.Follow(1, 3)
		got := tw.GetNewsFeed(1)
		want := []int{30, 20, 10}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("GetNewsFeed(1) = %v, want %v", got, want)
		}
	})
}
