# Get organization repositories

This returns list of repositories by passed organization.

## Inputs

### `github-organization`

**Required** The GitHub organization.

### `name-pattern`

**Optional** The regex pattern to match repository names. By default `.*`.

## Outputs

### `repositories`

The list of repositories related to passed organization.

## Example usage
```yaml
uses: actions/get-organization-repositoties@v1
with:
  github-organization: 'github'
```