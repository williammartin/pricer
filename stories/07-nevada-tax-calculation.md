# Nevada Tax Calculation

## Description

As a user placing orders in Nevada, I need the tool to calculate and include Nevada state tax so that I can provide accurate final pricing to my Nevada customers. The tool should apply Nevada's 8.00% tax rate to orders with state code "NV".

## Value

This continues expanding state tax support, allowing users to test another tax rate and ensure the tool correctly handles different state codes.

## Acceptance Criteria

Given I have an order with 25 items at $50 each in Nevada (subtotal = $1,250, qualifies for 3% discount)
When I run `pricer 25 50.00 NV`
Then I see the output `1309.50`

**Calculation breakdown for verification:**
- Subtotal: 25 × $50.00 = $1,250.00
- Discount: $1,250.00 × 3% = $37.50
- After discount: $1,250.00 - $37.50 = $1,212.50
- Tax: $1,212.50 × 8.00% = $97.00
- Total: $1,212.50 + $97.00 = $1,309.50