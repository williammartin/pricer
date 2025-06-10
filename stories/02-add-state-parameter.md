# Add State Code Parameter

## Description

As a user, I need to provide a state code along with my order details so that the tool can eventually calculate state-specific taxes. The pricer tool should accept the state code as a third parameter but continue to output just the subtotal for now.

## Value

This establishes the complete command-line interface that users will ultimately use, allowing them to provide feedback on the user experience and parameter order before we add tax calculation complexity.

## Acceptance Criteria

Given I have an order with 10 items at $50 each in Utah
When I run `pricer 10 50.00 UT`
Then I see the output `500.00`