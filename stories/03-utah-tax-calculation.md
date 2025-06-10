# Utah Tax Calculation

## Description

As a user placing orders in Utah, I need the tool to calculate and include Utah state tax so that I can provide accurate final pricing to my customers. The tool should apply Utah's 6.85% tax rate to orders with state code "UT".

## Value

This provides users with complete pricing information for Utah orders, allowing them to see the tax calculation in action and provide feedback on the accuracy and output format before we support additional states.

## Acceptance Criteria

Given I have an order with 10 items at $50 each in Utah
When I run `pricer 10 50.00 UT`
Then I see the output `534.25`

**Calculation breakdown for verification:**
- Subtotal: 10 × $50.00 = $500.00
- Tax: $500.00 × 6.85% = $34.25
- Total: $500.00 + $34.25 = $534.25