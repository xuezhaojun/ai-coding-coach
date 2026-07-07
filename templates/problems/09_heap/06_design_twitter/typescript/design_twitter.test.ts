import { describe, it, expect } from "vitest";
import { Twitter } from "./design_twitter";

describe("Twitter", () => {
  it("basic post and feed", () => {
    const tw = new Twitter();
    tw.postTweet(1, 5);
    expect(tw.getNewsFeed(1)).toEqual([5]);
  });

  it("follow and see followee tweets", () => {
    const tw = new Twitter();
    tw.postTweet(1, 5);
    tw.postTweet(2, 6);
    tw.follow(1, 2);
    expect(tw.getNewsFeed(1)).toEqual([6, 5]);
  });

  it("unfollow removes tweets from feed", () => {
    const tw = new Twitter();
    tw.postTweet(1, 5);
    tw.postTweet(2, 6);
    tw.follow(1, 2);
    tw.unfollow(1, 2);
    expect(tw.getNewsFeed(1)).toEqual([5]);
  });

  it("feed limited to 10 most recent", () => {
    const tw = new Twitter();
    for (let i = 1; i <= 12; i++) {
      tw.postTweet(1, i);
    }
    expect(tw.getNewsFeed(1)).toEqual([12, 11, 10, 9, 8, 7, 6, 5, 4, 3]);
  });

  it("empty feed for new user", () => {
    const tw = new Twitter();
    expect(tw.getNewsFeed(1)).toEqual([]);
  });

  it("multiple followees merged feed", () => {
    const tw = new Twitter();
    tw.postTweet(2, 10);
    tw.postTweet(3, 20);
    tw.postTweet(2, 30);
    tw.follow(1, 2);
    tw.follow(1, 3);
    expect(tw.getNewsFeed(1)).toEqual([30, 20, 10]);
  });
});
