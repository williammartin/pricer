# Pricer Specification Clarifications

This document captures the ambiguities from the original Pricer specification that were resolved during user story creation.

## Discount Application Logic

### Threshold Behavior
- **Ambiguity**: The original specification showed discount thresholds but didn't clarify if they were "at or above" vs "above only"
- **Clarification**: Discounts apply when the subtotal is **at or equal to** the threshold amount
  - Orders of exactly $1,000 receive 3% discount
  - Orders of exactly $5,000 receive 5% discount (not 3%)
  - And so on for all threshold values

### Discount Selection Logic  
- **Ambiguity**: When multiple thresholds are met, which discount applies?
- **Clarification**: The **highest applicable discount** is used based on the subtotal
  - A $6,000 order gets 5% discount (not 3%)
  - A $50,000 order gets 15% discount (not any lower tier)

### Order of Operations
- **Ambiguity**: Whether tax is calculated on original subtotal or after discount
- **Clarification**: Tax is calculated **after** discount is applied
  1. Calculate subtotal (quantity × price)
  2. Apply discount to subtotal
  3. Calculate tax on discounted amount
  4. Add tax to discounted amount for final total

## Input Validation Requirements

### Numeric Input Handling
- **Ambiguity**: How to handle non-numeric inputs
- **Clarification**: 
  - Non-numeric quantity: Show "Quantity must be a number"
  - Non-numeric price: Show "Price must be a number"
  - Negative quantity: Show "Quantity must be positive"
  - Negative price: Show "Price must be positive"

### State Code Validation
- **Ambiguity**: How to handle unsupported state codes
- **Clarification**: Show "State code not supported. Supported states: UT, NV, TX, AL, CA"

## Output Format

### Default Output
- **Ambiguity**: How detailed should the default output be?
- **Clarification**: Default output shows only the final total as a decimal number (e.g., `1295.56`)

### Verbose Mode
- **Ambiguity**: Should there be a way to see calculation breakdown?
- **Clarification**: A `--verbose` flag provides detailed breakdown:
  ```
  Subtotal: $1,250.00
  Discount (3%): -$37.50
  After discount: $1,212.50
  Tax (6.85%): $83.06
  Total: $1,295.56
  ```

## Supported States and Tax Rates
- **Clarification**: Only the five states listed in the specification table are supported:
  - UT: 6.85%
  - NV: 8.00% 
  - TX: 6.25%
  - AL: 4.00%
  - CA: 8.25%

## Command Line Interface
- **Clarification**: Tool is called `pricer` with three required arguments:
  - `pricer <quantity> <price> <state_code>`
  - Optional `--verbose` flag for detailed output