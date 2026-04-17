# Design Twitter

- **Difficulty**: Medium
- **Category**: Heap
- **Topics**: max-heap, hash map, design, merge k sorted lists

## Description

Design a simplified Twitter where users can post tweets, follow/unfollow other users, and retrieve the 10 most recent tweets in their news feed (from themselves and their followees).

## Approach

Store tweets per user with timestamps and followee sets with hash maps. For getNewsFeed, use a max-heap merge of k sorted lists (one per followed user plus self), extracting the 10 most recent tweets by timestamp.
