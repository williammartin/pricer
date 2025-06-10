# Non-Numeric Quantity Error Handling

## Description

As a user, I need a clear error message when I provide a non-numeric quantity value so that I understand the input format expected and can correct my mistake.

## Value

This helps users understand the correct input format when they accidentally provide text instead of numbers, improving the overall user experience.

## Acceptance Criteria

Given I provide a non-numeric quantity
When I run `pricer abc 50.00 UT`
Then I see the message "Quantity must be a number"