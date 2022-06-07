# `vish` Come True

## Installation

> Depends on the [`go` tool](https://go.dev/).

```bash
go install github.com/sharpvik/vish
mv $(go env GOPATH)/bin/vish /usr/local/bin/
```

## Make Default

> On Mac, users should list `/usr/local/bin/vish` in the `/etc/shells` file (try `sudo vim /etc/shells`) first; otherwise, `chsh` command will error out saying "non-standard shell".

```bash
chsh -s /usr/local/bin/vish
```
