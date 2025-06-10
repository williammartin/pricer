# Negative Price Error Handling

## Description

As a user, I need a clear error message when I provide a negative price value so that I understand that only positive prices are allowed for pricing calculations.

## Value

This prevents users from accidentally calculating prices with negative values, ensuring the tool only processes valid business pricing scenarios.

## Acceptance Criteria

Given I provide a negative price
When I run `pricer 10 -50.00 UT`
Then I see the message "Price must be positive"