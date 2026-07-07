interface Tweet {
  time: number;
  tweetId: number;
}

interface FeedItem {
  tweet: Tweet;
  userId: number;
  idx: number;
}

// Max-heap by tweet time.
class TweetHeap {
  private data: FeedItem[] = [];

  get length(): number {
    return this.data.length;
  }

  push(item: FeedItem): void {
    this.data.push(item);
    this.siftUp(this.data.length - 1);
  }

  pop(): FeedItem {
    const top = this.data[0];
    const last = this.data.pop()!;
    if (this.data.length > 0) {
      this.data[0] = last;
      this.siftDown(0);
    }
    return top;
  }

  private siftUp(i: number): void {
    while (i > 0) {
      const parent = Math.floor((i - 1) / 2);
      if (this.data[parent].tweet.time >= this.data[i].tweet.time) break;
      [this.data[parent], this.data[i]] = [this.data[i], this.data[parent]];
      i = parent;
    }
  }

  private siftDown(i: number): void {
    const n = this.data.length;
    while (true) {
      const left = 2 * i + 1;
      const right = 2 * i + 2;
      let largest = i;
      if (
        left < n &&
        this.data[left].tweet.time > this.data[largest].tweet.time
      )
        largest = left;
      if (
        right < n &&
        this.data[right].tweet.time > this.data[largest].tweet.time
      )
        largest = right;
      if (largest === i) break;
      [this.data[largest], this.data[i]] = [this.data[i], this.data[largest]];
      i = largest;
    }
  }
}

// Simplified Twitter implementation.
// PostTweet: O(1), GetNewsFeed: O(k log f) where k=10, f=#followees.
// Follow/Unfollow: O(1). Space: O(users + tweets).
export class SolveTwitter {
  private time = 0;
  private tweets: Map<number, Tweet[]> = new Map();
  private following: Map<number, Set<number>> = new Map();

  postTweet(userId: number, tweetId: number): void {
    this.time++;
    if (!this.tweets.has(userId)) this.tweets.set(userId, []);
    this.tweets.get(userId)!.push({ time: this.time, tweetId });
  }

  getNewsFeed(userId: number): number[] {
    const users: number[] = [userId];
    const followees = this.following.get(userId);
    if (followees) {
      for (const fid of followees) users.push(fid);
    }
    const heap = new TweetHeap();
    for (const uid of users) {
      const tweets = this.tweets.get(uid);
      if (tweets && tweets.length > 0) {
        const idx = tweets.length - 1;
        heap.push({ tweet: tweets[idx], userId: uid, idx });
      }
    }
    const result: number[] = [];
    while (heap.length > 0 && result.length < 10) {
      const item = heap.pop();
      result.push(item.tweet.tweetId);
      if (item.idx > 0) {
        const next = item.idx - 1;
        const tweet = this.tweets.get(item.userId)![next];
        heap.push({ tweet, userId: item.userId, idx: next });
      }
    }
    return result;
  }

  follow(followerId: number, followeeId: number): void {
    if (!this.following.has(followerId))
      this.following.set(followerId, new Set());
    this.following.get(followerId)!.add(followeeId);
  }

  unfollow(followerId: number, followeeId: number): void {
    this.following.get(followerId)?.delete(followeeId);
  }
}
