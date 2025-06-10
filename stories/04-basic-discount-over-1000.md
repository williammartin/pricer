# Basic Discount for Orders Over $1,000

## Description

As a user with larger orders, I need to receive a discount when my subtotal exceeds $1,000 so that I can offer competitive pricing to my customers. The tool should apply a 3% discount to the subtotal before calculating tax for orders over $1,000.

## Value

This introduces the discount functionality that users can test and provide feedback on before we add more complex discount tiers. Users with orders over $1,000 will immediately see value from reduced pricing.

## Acceptance Criteria

Given I have an order with 25 items at $50 each in Utah (subtotal = $1,250)
When I run `pricer 25 50.00 UT`
Then I see the output `1295.56`

**Calculation breakdown for verification:**
- Subtotal: 25 × $50.00 = $1,250.00
- Discount: $1,250.00 × 3% = $37.50
- After discount: $1,250.00 - $37.50 = $1,212.50
- Tax: $1,212.50 × 6.85% = $83.06
- Total: $1,212.50 + $83.06 = $1,295.56