# 5% Discount for Orders Over $5,000

## Description

As a user with high-value orders, I need to receive a larger discount when my subtotal exceeds $5,000 so that I can offer even more competitive pricing on bulk orders. The tool should apply a 5% discount to the subtotal before calculating tax for orders over $5,000.

## Value

This introduces the next discount tier, allowing users with larger orders to see increased savings and test the logic for selecting the appropriate discount rate based on order value.

## Acceptance Criteria

Given I have an order with 120 items at $50 each in Utah (subtotal = $6,000, qualifies for 5% discount)
When I run `pricer 120 50.00 UT`
Then I see the output `6087.75`

**Calculation breakdown for verification:**
- Subtotal: 120 × $50.00 = $6,000.00
- Discount: $6,000.00 × 5% = $300.00
- After discount: $6,000.00 - $300.00 = $5,700.00
- Tax: $5,700.00 × 6.85% = $390.45
- Total: $5,700.00 + $390.45 = $6,090.45