# Get organization repositories

This action returns a list of repositories related to passed organization. The action can be used with the matrix GitHub action feature to run jobs for a few repositories.

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
uses: denis-tingaikin/get-organization-repositories@v1
with:
  github-organization: 'github'
```
