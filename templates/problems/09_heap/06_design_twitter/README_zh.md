# 设计推特

- **难度**: 中等
- **分类**: 堆
- **考点**: 最大堆, 哈希表, 设计, 合并 K 个有序链表
- **链接**: [NeetCode](https://neetcode.io/problems/design-twitter-feed) | [力扣 355](https://leetcode.cn/problems/design-twitter/)

## 题目描述

设计一个简化版的推特，用户可以发布推文、关注/取消关注其他用户，以及获取其动态流中最近的 10 条推文 ID。实现以下操作：

- `ConstructorTwitter()` 初始化推特对象。
- `PostTweet(userId, tweetId int)` 用户 `userId` 发布一条推文，推文 ID 为 `tweetId`。每次调用此函数时 `tweetId` 唯一。
- `GetNewsFeed(userId int) []int` 检索用户动态流中最近的 10 条推文 ID。动态流中的每条推文必须是用户自己发布的或者其关注的用户发布的。推文必须按从最近到最久的顺序排列。
- `Follow(followerId, followeeId int)` ID 为 `followerId` 的用户开始关注 ID 为 `followeeId` 的用户。
- `Unfollow(followerId, followeeId int)` ID 为 `followerId` 的用户取消关注 ID 为 `followeeId` 的用户。

## 示例

**示例 1:**

```
输入:
  ConstructorTwitter()
  PostTweet(1, 5)
  GetNewsFeed(1)          -> [5]

解释: 用户 1 发布推文 5。其动态流中只包含自己的推文。
```

**示例 2:**

```
输入:
  ConstructorTwitter()
  PostTweet(1, 5)
  PostTweet(2, 6)
  Follow(1, 2)
  GetNewsFeed(1)          -> [6, 5]
  Unfollow(1, 2)
  GetNewsFeed(1)          -> [5]

解释: 用户 1 关注用户 2 后，其动态流显示两条推文（最新的在前）。取消关注后，只显示自己的推文。
```

**示例 3:**

```
输入:
  ConstructorTwitter()
  PostTweet(2, 10)
  PostTweet(3, 20)
  PostTweet(2, 30)
  Follow(1, 2)
  Follow(1, 3)
  GetNewsFeed(1)          -> [30, 20, 10]

解释: 用户 1 关注了用户 2 和 3。动态流按时间倒序合并他们的推文。
```

## 约束条件

- `1 <= userId, followerId, followeeId <= 500`
- `0 <= tweetId <= 10^4`
- 所有推文 ID 唯一。
- `PostTweet`、`GetNewsFeed`、`Follow` 和 `Unfollow` 的调用次数总计不超过 `3 * 10^4` 次。

## 函数签名

```go
type Twitter struct {}

func ConstructorTwitter() Twitter

func (t *Twitter) PostTweet(userId int, tweetId int)

func (t *Twitter) GetNewsFeed(userId int) []int

func (t *Twitter) Follow(followerId int, followeeId int)

func (t *Twitter) Unfollow(followerId int, followeeId int)
```
