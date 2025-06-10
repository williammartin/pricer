# Non-Numeric Price Error Handling

## Description

As a user, I need a clear error message when I provide a non-numeric price value so that I understand the input format expected for pricing calculations.

## Value

This helps users understand the correct price input format when they accidentally provide text instead of numbers, preventing calculation errors.

## Acceptance Criteria

Given I provide a non-numeric price
When I run `pricer 10 abc UT`
Then I see the message "Price must be a number"