# Negative Quantity Error Handling

## Description

As a user, I need a clear error message when I provide a negative quantity value so that I understand that only positive quantities are allowed.

## Value

This prevents users from accidentally calculating prices with negative quantities, ensuring the tool only processes valid business scenarios.

## Acceptance Criteria

Given I provide a negative quantity
When I run `pricer -5 50.00 UT`
Then I see the message "Quantity must be positive"