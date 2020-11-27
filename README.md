# Get organization repositories

This returns a list of repositories by the passed organizations. The action can be used with the matrix GitHub action feature to run jobs for a few repositories.

## Inputs

### `github-organization`

**Required** The GitHub organization.

### `name-pattern`

**Optional** The regex pattern to match repository names. By default `.*`.

## Outputs

### `repositories`

The list of repositories related to passed organization in JSON format.

## Example usage
```yaml
uses: actions/get-organization-repositoties@v1
with:
  github-organization: 'github'
```