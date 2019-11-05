# semversort
Sort semantic version numbers from stdin

## Install

With go 1.13+

```bash
go get github.com/whereswaldon/semversort@latest
```

If you encounter an error about module support, `export GO111MODULE=on` and try again.

## Use

Example use:

```bash
curl -sS https://api.github.com/repos/kubernetes/kubernetes/releases \
  | jq -r '.[].tag_name' | semversort
```
