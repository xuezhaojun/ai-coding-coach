from design_twitter import Twitter


def test_basic_post_and_feed():
    tw = Twitter()
    tw.post_tweet(1, 5)
    assert tw.get_news_feed(1) == [5]


def test_follow_and_see_followee_tweets():
    tw = Twitter()
    tw.post_tweet(1, 5)
    tw.post_tweet(2, 6)
    tw.follow(1, 2)
    assert tw.get_news_feed(1) == [6, 5]


def test_unfollow_removes_tweets_from_feed():
    tw = Twitter()
    tw.post_tweet(1, 5)
    tw.post_tweet(2, 6)
    tw.follow(1, 2)
    tw.unfollow(1, 2)
    assert tw.get_news_feed(1) == [5]


def test_feed_limited_to_10_most_recent():
    tw = Twitter()
    for i in range(1, 13):
        tw.post_tweet(1, i)
    assert tw.get_news_feed(1) == [12, 11, 10, 9, 8, 7, 6, 5, 4, 3]


def test_empty_feed_for_new_user():
    tw = Twitter()
    assert tw.get_news_feed(1) == []


def test_multiple_followees_merged_feed():
    tw = Twitter()
    tw.post_tweet(2, 10)
    tw.post_tweet(3, 20)
    tw.post_tweet(2, 30)
    tw.follow(1, 2)
    tw.follow(1, 3)
    assert tw.get_news_feed(1) == [30, 20, 10]
