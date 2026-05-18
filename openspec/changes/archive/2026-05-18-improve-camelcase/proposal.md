## Why

The current `CamelCase` function only has a stub test that calls `CamelCase("a")` without any assertions. The function handles multiple input formats (spaces, hyphens, snake_case, CONSTANT_CASE, PascalCase) but has no automated verification that any of these work correctly.

## What Changes

- Add comprehensive test cases covering all documented input formats:
  - Space-separated strings (`hello world`)
  - Hyphen-separated strings (`some-hyphen-text`)
  - CONSTANT_CASE strings (`CONSTANT_CASE`)
  - PascalCase strings (`PascalCase`)
  - Mixed case with multiple spaces (`mixed   SpAcE`)
- Verify the implementation handles these correctly; fix any bugs uncovered by the tests

## Capabilities

### New Capabilities

- `camelcase-format-conversion`: Verify `CamelCase` correctly converts space-separated, hyphen-separated, snake_case with all caps, PascalCase, and mixed-case inputs to camelCase

### Modified Capabilities

<!-- None -->

## Impact

- Affected code: `stringkit/camel_case.go`, `stringkit/camel_case_test.go`
- No API or dependency changes
