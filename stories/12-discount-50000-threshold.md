# 15% Discount for Orders Over $50,000

## Description

As a user with very large enterprise orders, I need to receive the maximum discount when my subtotal exceeds $50,000 so that I can offer the best possible pricing on massive bulk purchases. The tool should apply a 15% discount to the subtotal before calculating tax for orders over $50,000.

## Value

This completes the discount tier system with the highest discount rate, allowing users to test the full range of discount calculations and ensure maximum savings for their largest orders.

## Acceptance Criteria

Given I have an order with 1100 items at $50 each in Utah (subtotal = $55,000, qualifies for 15% discount)
When I run `pricer 1100 50.00 UT`
Then I see the output `49702.88`

**Calculation breakdown for verification:**
- Subtotal: 1,100 × $50.00 = $55,000.00
- Discount: $55,000.00 × 15% = $8,250.00
- After discount: $55,000.00 - $8,250.00 = $46,750.00
- Tax: $46,750.00 × 6.85% = $3,202.38
- Total: $46,750.00 + $3,202.38 = $49,952.38