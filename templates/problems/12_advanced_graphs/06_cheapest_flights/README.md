# Cheapest Flights Within K Stops

- **Difficulty**: Medium
- **Category**: Advanced Graphs
- **Topics**: Bellman-Ford, dynamic programming, shortest path

## Description

Given n cities, flights with prices, a source, destination, and maximum number of stops k, find the cheapest price from src to dst with at most k stops. Return -1 if no such route exists.

## Approach

Use a modified Bellman-Ford algorithm running k+1 iterations (k stops = k+1 edges). In each iteration, relax all edges using the prices from the previous iteration (copy array to prevent using updates from the current round). This ensures we only consider paths with at most k+1 edges.
