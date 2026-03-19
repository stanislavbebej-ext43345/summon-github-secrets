[![Build](https://github.com/stanislavbebej-ext43345/summon-github-secrets/actions/workflows/build.yml/badge.svg)](.github/workflows/build.yml)
[![dependabot](https://img.shields.io/badge/Dependabot-enabled-brightgreen?logo=dependabot)](.github/dependabot.yml)
[![editorconfig](https://img.shields.io/badge/EditorConfig-enabled-brightgreen?logo=editorconfig)](.editorconfig)
[![release-please](https://img.shields.io/badge/release--please-enabled-brightgreen?logo=google)](release-please-config.json)

# summon-github-secrets

[GitHub Actions Secrets](https://docs.github.com/en/actions/security-for-github-actions/security-guides/using-secrets-in-github-actions) provider for [Summon](https://github.com/cyberark/summon) in go.

## Development

```bash
export BINARY_NAME="summon-github-secrets"

go build -ldflags "-s -w" -o $BINARY_NAME
strip $BINARY_NAME
upx -q -9 $BINARY_NAME

sudo cp $BINARY_NAME /usr/local/bin/Providers
```

## Usage

1. create a [secrets.yml](./secrets.yml) configuration file with `secretId`s:

```yaml
SECRET_VARIABLE: !var github_token
```

2. run `summon`:

```bash
export GITHUB_SECRETS='{"github_token":"***"}'

summon -p summon-github-secrets cat @SUMMONENVFILE
```
