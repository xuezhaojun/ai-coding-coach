import heapq


class SolveTwitter:
    """Simplified Twitter implementation.

    PostTweet: O(1), GetNewsFeed: O(k log f) where k=10, f=#followees.
    Follow/Unfollow: O(1). Space: O(users + tweets).
    """

    def __init__(self) -> None:
        self._time = 0
        self._tweets: dict[int, list[tuple[int, int]]] = {}
        self._following: dict[int, set[int]] = {}

    def post_tweet(self, user_id: int, tweet_id: int) -> None:
        self._time += 1
        self._tweets.setdefault(user_id, []).append((self._time, tweet_id))

    def get_news_feed(self, user_id: int) -> list[int]:
        # Include self.
        users = [user_id]
        users.extend(self._following.get(user_id, set()))
        # Max-heap by timestamp (negate).
        heap: list[tuple[int, int, int, int]] = []
        for uid in users:
            tweets = self._tweets.get(uid, [])
            if tweets:
                idx = len(tweets) - 1
                heapq.heappush(heap, (-tweets[idx][0], uid, idx, tweets[idx][1]))
        result: list[int] = []
        while heap and len(result) < 10:
            _, uid, idx, tweet_id = heapq.heappop(heap)
            result.append(tweet_id)
            if idx > 0:
                nxt = idx - 1
                t = self._tweets[uid][nxt]
                heapq.heappush(heap, (-t[0], uid, nxt, t[1]))
        return result

    def follow(self, follower_id: int, followee_id: int) -> None:
        self._following.setdefault(follower_id, set()).add(followee_id)

    def unfollow(self, follower_id: int, followee_id: int) -> None:
        if follower_id in self._following:
            self._following[follower_id].discard(followee_id)
