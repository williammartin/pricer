# 7% Discount for Orders Over $7,000

## Description

As a user with very high-value orders, I need to receive an even larger discount when my subtotal exceeds $7,000 so that I can maximize savings on substantial bulk orders. The tool should apply a 7% discount to the subtotal before calculating tax for orders over $7,000.

## Value

This adds another discount tier, allowing users to test the progression of discount rates and ensuring the tool selects the correct discount based on order value thresholds.

## Acceptance Criteria

Given I have an order with 150 items at $50 each in Utah (subtotal = $7,500, qualifies for 7% discount)
When I run `pricer 150 50.00 UT`
Then I see the output `7451.44`

**Calculation breakdown for verification:**
- Subtotal: 150 × $50.00 = $7,500.00
- Discount: $7,500.00 × 7% = $525.00
- After discount: $7,500.00 - $525.00 = $6,975.00
- Tax: $6,975.00 × 6.85% = $477.79
- Total: $6,975.00 + $477.79 = $7,452.79