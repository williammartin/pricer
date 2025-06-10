# Basic Subtotal Calculation

## Description

As a user, I need to calculate the subtotal for an order so that I can quickly determine the cost before any additional charges. The pricer tool should accept the number of items and price per item as command-line arguments and output the subtotal.

## Value

This provides the foundational calculation that users need to see immediate results and verify the basic multiplication logic works correctly, allowing them to provide feedback on the core functionality.

## Acceptance Criteria

Given I have an order with 10 items at $50 each
When I run `pricer 10 50.00`
Then I see the output `500.00`