export class Twitter {
  constructor() {}

  postTweet(userId: number, tweetId: number): void {}

  getNewsFeed(userId: number): number[] {
    return [];
  }

  follow(followerId: number, followeeId: number): void {}

  unfollow(followerId: number, followeeId: number): void {}
}
