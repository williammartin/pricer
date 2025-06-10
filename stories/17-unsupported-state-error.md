# Unsupported State Code Error Handling

## Description

As a user, I need a clear error message when I provide a state code that isn't supported so that I understand which states the tool can calculate taxes for.

## Value

This helps users understand the current limitations of the tool and which state codes are valid, preventing confusion when they try to use unsupported states.

## Acceptance Criteria

Given I provide an unsupported state code
When I run `pricer 10 50.00 NY`
Then I see the message "State code not supported. Supported states: UT, NV, TX, AL, CA"