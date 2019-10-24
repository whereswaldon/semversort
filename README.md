# semversort
Sort semantic version numbers from stdin

## Install

With go 1.13+

```bash
go get github.com/whereswaldon/semversort@latest
```

## Use

Example use:

```bash
curl https://api.github.com/repos/ul/kak-lsp/releases | jq .[].tag_name | tr -d '"' | shuf | semversort
```

Note that the output would have been sorted by github without the `shuf` in the middle there.
