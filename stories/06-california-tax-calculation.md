# California Tax Calculation

## Description

As a user placing orders in California, I need the tool to calculate and include California state tax so that I can provide accurate final pricing to my California customers. The tool should apply California's 8.25% tax rate to orders with state code "CA".

## Value

This adds support for California's higher tax rate, allowing users to test how the tool handles the highest tax rate in our supported states and provide feedback on accuracy.

## Acceptance Criteria

Given I have an order with 25 items at $50 each in California (subtotal = $1,250, qualifies for 3% discount)
When I run `pricer 25 50.00 CA`
Then I see the output `1312.53`

**Calculation breakdown for verification:**
- Subtotal: 25 × $50.00 = $1,250.00
- Discount: $1,250.00 × 3% = $37.50
- After discount: $1,250.00 - $37.50 = $1,212.50
- Tax: $1,212.50 × 8.25% = $100.03
- Total: $1,212.50 + $100.03 = $1,312.53