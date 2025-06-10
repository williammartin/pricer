# 10% Discount for Orders Over $10,000

## Description

As a user with enterprise-level orders, I need to receive a substantial discount when my subtotal exceeds $10,000 so that I can offer highly competitive pricing on large-scale purchases. The tool should apply a 10% discount to the subtotal before calculating tax for orders over $10,000.

## Value

This introduces a significant discount milestone that provides substantial savings for high-value orders, allowing users to test the double-digit discount logic.

## Acceptance Criteria

Given I have an order with 250 items at $50 each in Utah (subtotal = $12,500, qualifies for 10% discount)
When I run `pricer 250 50.00 UT`
Then I see the output `12020.63`

**Calculation breakdown for verification:**
- Subtotal: 250 × $50.00 = $12,500.00
- Discount: $12,500.00 × 10% = $1,250.00
- After discount: $12,500.00 - $1,250.00 = $11,250.00
- Tax: $11,250.00 × 6.85% = $770.625 (rounds to $770.63)
- Total: $11,250.00 + $770.625 = $12,020.625 (rounds to $12,020.63)