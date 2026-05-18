## Context

The `CamelCase` function in `stringkit/camel_case.go` already implements the logic for converting multiple input formats to camelCase. The current test file only contains a stub with no assertions. The implementation appears to handle all documented cases correctly.

## Goals / Non-Goals

**Goals:**
- Add comprehensive table-driven tests covering all documented input formats
- Verify the existing implementation passes all tests
- Fix any bugs uncovered by the test suite

**Non-Goals:**
- Acronym handling (e.g., `JSONParser` → `jsonParser`)
- Unicode segmentation beyond basic ASCII / single-rune boundaries
- Other string conversion methods (PascalCase, snake_case, etc.)

## Decisions

- **Table-driven tests**: Use Go's standard table-driven test pattern for readability and easy extension
- **No implementation changes unless tests fail**: The current logic handles the required cases; only fix bugs if discovered

## Risks / Trade-offs

- If the existing implementation has edge-case bugs not covered by the stated test cases, they will remain — acceptable since the scope is limited to the documented cases
