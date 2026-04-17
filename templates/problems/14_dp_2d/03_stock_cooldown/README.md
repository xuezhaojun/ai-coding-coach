# Best Time to Buy and Sell Stock with Cooldown

- **Difficulty**: Medium
- **Category**: 2D Dynamic Programming
- **Topics**: dynamic programming, state machine

## Description

Given an array of stock prices, find the maximum profit with unlimited transactions but a one-day cooldown after selling (you cannot buy the day after you sell).

## Approach

Model three states: held (holding a stock), sold (just sold today), and rest (idle/cooldown). Transition: sold = held + price, held = max(held, rest - price), rest = max(rest, prevSold). Track these with three variables for O(1) space.
