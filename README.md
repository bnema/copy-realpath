# crp (copy-realpath)

A simple CLI tool that copies file paths to your clipboard.

## Install

```bash
go install github.com/bnema/copy-realpath/cmd/crp@latest
```

## Usage

```bash
crp              # copies current directory
crp .            # copies current directory  
crp file.txt     # copies absolute path to file.txt
crp ../foo       # copies resolved absolute path
```

## Requirements

Requires `wl-copy` (Wayland) or `xclip` (X11) to be installed on your system.
