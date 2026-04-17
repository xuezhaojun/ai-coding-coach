package design_twitter

type Twitter struct {
}

func ConstructorTwitter() Twitter {
	return Twitter{}
}

func (t *Twitter) PostTweet(userId int, tweetId int) {
}

func (t *Twitter) GetNewsFeed(userId int) []int {
	return nil
}

func (t *Twitter) Follow(followerId int, followeeId int) {
}

func (t *Twitter) Unfollow(followerId int, followeeId int) {
}
