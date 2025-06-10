# Texas Tax Calculation

## Description

As a user placing orders in Texas, I need the tool to calculate and include Texas state tax so that I can provide accurate final pricing to my Texas customers. The tool should apply Texas's 6.25% tax rate to orders with state code "TX".

## Value

This expands tax calculation support to a second state, allowing users to test the state-specific logic and provide feedback on how the tool handles different tax rates.

## Acceptance Criteria

Given I have an order with 25 items at $50 each in Texas (subtotal = $1,250, qualifies for 3% discount)
When I run `pricer 25 50.00 TX`
Then I see the output `1288.31`

**Calculation breakdown for verification:**
- Subtotal: 25 × $50.00 = $1,250.00
- Discount: $1,250.00 × 3% = $37.50
- After discount: $1,250.00 - $37.50 = $1,212.50
- Tax: $1,212.50 × 6.25% = $75.78
- Total: $1,212.50 + $75.78 = $1,288.28