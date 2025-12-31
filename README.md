# cpt (copy-that-path)

A simple CLI tool that copies file paths to your clipboard.

## Install

```bash
go install github.com/bnema/copy-that-path/cmd/cpt@latest
```

## Usage

```bash
cpt              # copies current directory
cpt .            # copies current directory  
cpt file.txt     # copies absolute path to file.txt
cpt ../foo       # copies resolved absolute path
```

## Requirements

Requires `wl-copy` (Wayland) or `xclip` (X11) to be installed on your system.
