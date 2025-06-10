# Alabama Tax Calculation

## Description

As a user placing orders in Alabama, I need the tool to calculate and include Alabama state tax so that I can provide accurate final pricing to my Alabama customers. The tool should apply Alabama's 4.00% tax rate to orders with state code "AL".

## Value

This completes tax calculation support for all required states, with Alabama having the lowest tax rate, allowing users to test the full range of tax calculations.

## Acceptance Criteria

Given I have an order with 25 items at $50 each in Alabama (subtotal = $1,250, qualifies for 3% discount)
When I run `pricer 25 50.00 AL`
Then I see the output `1261.00`

**Calculation breakdown for verification:**
- Subtotal: 25 × $50.00 = $1,250.00
- Discount: $1,250.00 × 3% = $37.50
- After discount: $1,250.00 - $37.50 = $1,212.50
- Tax: $1,212.50 × 4.00% = $48.50
- Total: $1,212.50 + $48.50 = $1,261.00