### Requirement: Space-separated strings
The system SHALL convert space-separated strings to camelCase by removing spaces and capitalizing the first letter of each word after the first.

#### Scenario: Basic space-separated input
- **WHEN** `CamelCase` is called with `"hello world"`
- **THEN** it returns `"helloWorld"`

#### Scenario: Multiple consecutive spaces
- **WHEN** `CamelCase` is called with `"hello   world"`
- **THEN** it returns `"helloWorld"`

### Requirement: Hyphen-separated strings
The system SHALL convert hyphen-separated strings to camelCase by removing hyphens and capitalizing the first letter of each word after the first.

#### Scenario: Basic hyphen input
- **WHEN** `CamelCase` is called with `"some-hyphen-text"`
- **THEN** it returns `"someHyphenText"`

### Requirement: CONSTANT_CASE strings
The system SHALL convert CONSTANT_CASE (all-uppercase with underscore separators) strings to camelCase by splitting on underscores, lowercasing all letters, and capitalizing the first letter of each word after the first.

#### Scenario: CONSTANT_CASE input
- **WHEN** `CamelCase` is called with `"CONSTANT_CASE"`
- **THEN** it returns `"constantCase"`

### Requirement: PascalCase strings
The system SHALL convert PascalCase strings to camelCase by lowercasing the first character and preserving subsequent word boundaries.

#### Scenario: PascalCase input
- **WHEN** `CamelCase` is called with `"PascalCase"`
- **THEN** it returns `"pascalCase"`

### Requirement: Mixed case with spaces
The system SHALL preserve the original casing of characters in mixed-case inputs, only removing spaces and capitalizing the first letter after a space boundary.

#### Scenario: Mixed case with multiple spaces
- **WHEN** `CamelCase` is called with `"mixed   SpAcE"`
- **THEN** it returns `"mixedSpAcE"`
