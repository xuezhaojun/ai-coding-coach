# Reconstruct Itinerary

- **Difficulty**: Hard
- **Category**: Advanced Graphs
- **Topics**: Eulerian path, DFS, Hierholzer's algorithm, greedy

## Description

Given a list of airline tickets represented as [from, to] pairs, reconstruct the itinerary starting from "JFK". If there are multiple valid itineraries, return the one with the smallest lexical order.

## Approach

Use Hierholzer's algorithm to find an Eulerian path. Build an adjacency list and sort destinations lexicographically. Use DFS, always visiting the smallest available destination first and removing the used edge. Append airports to the result in post-order, then reverse the result to get the correct itinerary.
