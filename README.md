[![Build Status](https://api.travis-ci.org/binaryplease/base16-universal-manager.svg)](http://travis-ci.org/binaryplease/base16-universal-manager) [![GoDoc](https://godoc.org/github.com/binaryplease/base16-universal-manager?status.svg)](http://godoc.org/github.com/binaryplease/base16-universal-manager)
[![Go Report Card](https://goreportcard.com/badge/github.com/binaryplease/base16-universal-manager)](https://goreportcard.com/report/github.com/binaryplease/base16-universal-manager)
[![codecov](https://codecov.io/gh/binaryplease/base16-universal-manager/branch/master/graph/badge.svg)](https://codecov.io/gh/binaryplease/base16-universal-manager)
[![Maintainability](https://api.codeclimate.com/v1/badges/65217f7940ee0e37d474/maintainability)](https://codeclimate.com/github/binaryplease/base16-universal-manager/maintainability)
[![HitCount](http://hits.dwyl.io/binaryplease/base16-universal-manager.svg)](http://hits.dwyl.io/binaryplease/base16-universal-manager)


# ![Base16](logo.png)

## About
### Base16 Universal Manager
Gets a [base16](https://github.com/chriskempson/base16) colorscheme the configured templates from the official repos and renders them out to the given locatinos
TODO

### Base16 themes
An architecture for building themes based on carefully chosen syntax
highlighting using a base of sixteen colors. Base16 provides a set of guidelines
detailing how to style syntax and how to code a _builder_ for compiling Base16
_schemes_ and _templates_.

### Why another manager/builder?
TODO
#### Similar projects

## Installation

At the moment, you can only build and install this with go. You can install it
directly or build it from source. I might provide packages for multiple linux
distributions and pre-build binaries in the future.

### Install directly
```
go install github.com/binaryplease/base16-universal-manager
```
### Get source and build manually
```
go get github.com/binaryplease/base16-universal-manager
cd $GOPATH/src/github.com/binaryplease/base16-universal-manager
go build
```

If you get errors on missing dependencies, install them as usual with `go get`.

## Usage

To run, just execute the application without any command line flags. It will
expect a config.yaml (example provided) in the same directory and render all
specified application templates with the selected colorscheme.

The following flags are planned and will be implemented soon:
```
usage: base16-setter [<flags>]

Flags:
  --help             Show context-sensitive help (also try --help-long and --help-man).
  --update-list      Update the list of templates and colorschemes
  --clear-list       Delete local master list caches
  --clear-templates  Delete local scheme caches
  --clear-schemes    Delete local template caches
  --version          Show application version.
```

## Configuration
### GitHub Token
TODO
### Configuration file
TODO

### Applications
Base16 Universal Manager can support all applications listed in the base16 repo.
For application-specific integration examples see the following list.

#### alacritty
#### binary-ninja
#### blink
#### c_header
#### concfg
#### conemu
#### console2
#### crosh
#### dunst
#### emacs
#### fzf
#### gnome-terminal
#### godot
#### gtk2
#### highlight
#### html-preview
#### i3
#### i3status-rust
#### iterm2
#### jetbrains
#### joe
#### kakoune
#### kitty
#### konsole
#### mintty
#### monodevelop
#### prism
#### prompt-toolkit
#### putty
#### pygments
#### pywal
#### qtcreator
#### qutebrowser
#### radare2
#### rofi
#### shell
#### scide
#### st
#### stumpwm
#### styles
#### telegram-desktop
#### termite
#### termux
#### textmate
#### tmux
#### tilix
#### vim
#### vis
#### vscode
#### windows-command-prompt
#### xcode
#### xfce4-terminal
#### xresources
#### xshell
#### zathura

## Contributing
I hacked this project together in a weekend and it grew to be bigger than
expected. The code quality could be way better and even though it is already
pretty usable, you might find bugs or other issues. The documentation is
work-in-progress.

Issues, bug-reports, pull requests or ideas for features and improvements are
**very welcome**. Also it would be great if users of specific applications can
document the usage of their respective templates, as I don't use all of them and
can't/won't test the integration for every single application.

