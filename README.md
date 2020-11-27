# Get organization repositories

This returns list of repositories by passed organization.

## Inputs

### `github-organization`

**Required** The GitHub organization.

## Outputs

The list of repositories related to passed organization.

## Example usage

uses: actions/get-organization-repositoties@v1
with:
  github-organization: 'github'