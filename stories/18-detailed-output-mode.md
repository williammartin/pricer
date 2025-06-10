# Detailed Output Mode

## Description

As a user, I need to see a detailed breakdown of pricing calculations so that I can verify the math and understand how the final price was determined. The tool should provide a verbose mode that shows subtotal, discount, tax, and final total.

## Value

This allows users to verify calculations, understand the pricing logic, and provide detailed pricing breakdowns to their customers when needed.

## Acceptance Criteria

Given I have an order with 25 items at $50 each in Utah (qualifies for 3% discount)
When I run `pricer 25 50.00 UT --verbose`
Then I see the output:
```
Subtotal: $1,250.00
Discount (3%): -$37.50
After discount: $1,212.50
Tax (6.85%): $83.06
Total: $1,295.56
```