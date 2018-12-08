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
Base16 themes look great and the fact that you can use them on a lot of
different applications makes it easy to get a desktop configuration set up
easealy that looks uniformly themed. I really like the idea of having theme
colors and the templates for different applications separated.

The common workflow to set up a theme on your system looks like this in most
cases:

1. Search for a theme you like
2. Scroll through the support list of supported applications to find all the
   ones you need
3. Download or build the theme files for every application
4. Copy them all in their respective places.
5. Restart or refresh every application

I found this pretty tedious to do, since I like to change my theme every now and
then. This is where the need for a manager came up.

The idea: automate the complete workflow above. After setting up my
configuration for the manager **once** I ideally want to run **one** command and
have all my applicatins set-up and ready whith the theme I choose.

There are other projects that try to simplify this workflow, but I have had
problems with them in the past and don't particulary like the designs. Setting a
theme should *not* require me to download or build them *all*. It also should
really set them, avoiding having to run a theme-setting script every time the
application is started.

### What Base16 Universal Manager is NOT
Even though this project builds the themes needed and could probably be expanded
or used as a theme builder in the sense described in the official base16
guidelines, this is not what it is intended for. It is mainly aimed at users and
not theme or template maintainers and desigend to only get and build the stuff
the user really needs.

#### Similar projects
[base16-shell]() A shell script to change your shell's default ANSI colors

I liked the idea, but it limits the use to command line applications.
Also I found the script to be to slow on my system, which results in the colors
of new terminals been changed about ~0.5s after start. In the meantime my
terminal waits. There probably are usecases, where this is the better choice,
but it was not what I was looking for.

[base16-manager]() A command line tool to install base16 templates and set themes globally.

Even though this project aimed to provide the same funcionality, only very few
applications are supported. Also it required me to download a lot of repos I
will never use.


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
### GitHub Token (optional)
This program fetches data using the GitHub API. If you are not logged in, you
might get blocked by the API's [rate limiting](https://developer.github.com/v3/#rate-limiting)

To increase the amout of allowed request, you can use personal access-token.
Generate one [here](https://github.com/settings/tokens/new) (Default options
should be enough, just provide a name) and put it in your configuration file.

Putting the token in the configuration will automatically use it to make all
requests as a registered user.

### Configuration file

The configuration file specifies everything the program should do when run. It
consists mainly of two parts:

#### General configuration values

| Variable | Default | Explanation|
| ---|---|---|
| GithubToken |  "set-your-token-here" | see `GitHub Token (optinonal)`|
| Colorscheme |  "flat.yml" | The colorscheme to use |
| SchemesListFile |  "cache/schemeslist.yaml" | cache file for the list of
Colorschemes|
| TemplatesListFile |  "cache/templateslist.yaml" | Cache file for the list of
templates |
| SchemesCachePath |  "cache/schemes/" | Colorschemes cache directory| 
| TemplatesCachePath |  "cache/templates/" | Templates cache directory |
| DryRun |  false | Print the rendered files to stdout instead of saving them|


#### Applications which you want to theme



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

