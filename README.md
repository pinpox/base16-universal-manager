[![Build Status](https://api.travis-ci.org/binaryplease/base16-universal-manager.svg)](http://travis-ci.org/binaryplease/base16-universal-manager) [![GoDoc](https://godoc.org/github.com/binaryplease/base16-universal-manager?status.svg)](http://godoc.org/github.com/binaryplease/base16-universal-manager)
[![Go Report Card](https://goreportcard.com/badge/github.com/binaryplease/base16-universal-manager)](https://goreportcard.com/report/github.com/binaryplease/base16-universal-manager)
[![codecov](https://codecov.io/gh/binaryplease/base16-universal-manager/branch/master/graph/badge.svg)](https://codecov.io/gh/binaryplease/base16-universal-manager)
[![HitCount](http://hits.dwyl.io/binaryplease/base16-universal-manager.svg)](http://hits.dwyl.io/binaryplease/base16-universal-manager)


# ![Base16](logo.png)

## About
### Base16 Universal Manager
Gets a [base16](https://github.com/chriskempson/base16) colorscheme the configured templates from the official repos and renders them out to the given locatinos
TODO
TODO: comment about levenstein?

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
| GithubToken |  set-your-token-here | see `GitHub Token (optinonal)`|
| Colorscheme |  flat.yml | The colorscheme to use |
| SchemesListFile |  cache/schemeslist.yaml | cache file for the list of Colorschemes|
| TemplatesListFile |  cache/templateslist.yaml | Cache file for the list of templates |
| SchemesCachePath |  cache/schemes/ | Colorschemes cache directory| 
| TemplatesCachePath |  cache/templates/ | Templates cache directory |
| DryRun |  false | Print the rendered files to stdout instead of saving them|

The provided theme names, do not have to be exact.
The [Levenstein distance](https://en.wikipedia.org/wiki/Levenshtein_distance<Paste>)
is used to calculate the best matching option. This is handy in case you only
partly remember the name of a particular scheme, if you made a typo, or if you are just plain *lazy*.




#### Applications which you want to theme
The rest of the configuration are the application specific settings. It consists
of a list of applicions you want to use. Here is an example that would set up
vim and i3 for you:

```
applications:
  i3:
    enabled: true
    hook: i3-msg 'restart'
    files:
      default: "~/.i3/i3_colors"
      bar-colors: "~/.i3/i3_bar_colors"
  vim:
    enabled: true
    files:
      default: "~/.vim/vim_colors"
```

In this configuration we render the files called `default` and `bar-colors` from
the [base16-i3 templates repository](https://github.com/khamer/base16-i3/tree/master/templates)
to a to `~/.i3/i3_colors` and `~/.i3/i3_bar_colors` as
well as the file called `default` from the [base16-vim templates
repo](https://github.com/chriskempson/base16-vim/tree/master/templates) to
`~/.vim/vim_colors`. In the configurations of those applications you could then
source that generated files.

The `hook` variable can be set for every application configured. It allows to
run a command after the files have been rendererd, e.g. to refersh the
application in order to show the new values. In this example we use `i3-msg
'restart'` to restart the i3 window manager in place, thus reloading the colors.

### Application integration examples
Base16 Universal Manager can support all applications listed in the base16 repo.
For application-specific integration examples see the following list.

#### alacritty
- Repo: https://github.com/aaron-williamson/base16-alacritty


Example configuration:
```

//TODO

```

#### binary-ninja
- Repo: https://github.com/evanrichter/base16-binary-ninja

Example configuration:
```

//TODO

```
#### blink
- Repo: https://github.com/niklaas/base16-blink.git

Example configuration:
```

//TODO

```
#### c_header
- Repo: https://github.com/m1sports20/base16-c_header

Example configuration:
```

//TODO

```
#### concfg
- Repo: https://github.com/h404bi/base16-concfg

Example configuration:
```

//TODO

```
#### conemu
- Repo: https://github.com/martinlindhe/base16-conemu

Example configuration:
```

//TODO

```
#### console2
- Repo: TODO

Example configuration:
```

//TODO

```
#### crosh
- Repo: TODO

Example configuration:
```

//TODO

```
#### dunst
- Repo: TODO

Example configuration:
```

//TODO

```
#### emacs
- Repo: TODO

Example configuration:
```

//TODO

```
#### fzf
- Repo: TODO

Example configuration:
```

//TODO

```
#### gnome-terminal
- Repo: TODO

Example configuration:
```

//TODO

```
#### godot
- Repo: TODO

Example configuration:
```

//TODO

```
#### gtk2
- Repo: TODO

Example configuration:
```

//TODO

```
#### highlight
- Repo: TODO

Example configuration:
```

//TODO

```
#### html-preview
- Repo: TODO

Example configuration:
```

//TODO

```
#### i3
- Repo: TODO

Example configuration:
```

//TODO

```
#### i3status-rust
- Repo: TODO

Example configuration:
```

//TODO

```
#### iterm2
- Repo: TODO

Example configuration:
```

//TODO

```
#### jetbrains
- Repo: TODO

Example configuration:
```

//TODO

```
#### joe
- Repo: TODO

Example configuration:
```

//TODO

```
#### kakoune
- Repo: TODO

Example configuration:
```

//TODO

```
#### kitty
- Repo: TODO

Example configuration:
```

//TODO

```
#### konsole
- Repo: TODO

Example configuration:
```

//TODO

```
#### mintty
- Repo: TODO

Example configuration:
```

//TODO

```
#### monodevelop
- Repo: TODO

Example configuration:
```

//TODO

```
#### prism
- Repo: TODO

Example configuration:
```

//TODO

```
#### prompt-toolkit
- Repo: TODO

Example configuration:
```

//TODO

```
#### putty
- Repo: TODO

Example configuration:
```

//TODO

```
#### pygments
- Repo: TODO

Example configuration:
```

//TODO

```
#### pywal
- Repo: TODO

Example configuration:
```

//TODO

```
#### qtcreator
- Repo: TODO

Example configuration:
```

//TODO

```
#### qutebrowser
- Repo: TODO

Example configuration:
```

//TODO

```
#### radare2
- Repo: TODO

Example configuration:
```

//TODO

```
#### rofi
- Repo: TODO

Example configuration:
```

//TODO

```
#### shell
- Repo: TODO

Example configuration:
```

//TODO

```
#### scide
- Repo: TODO

Example configuration:
```

//TODO

```
#### st
- Repo: TODO

Example configuration:
```

//TODO

```
#### stumpwm
- Repo: TODO

Example configuration:
```

//TODO

```
#### styles
- Repo: TODO

Example configuration:
```

//TODO

```
#### telegram-desktop
- Repo: TODO

Example configuration:
```

//TODO

```
#### termite
- Repo: TODO

Example configuration:
```

//TODO

```
#### termux
- Repo: TODO

Example configuration:
```

//TODO

```
#### textmate
- Repo: TODO

Example configuration:
```

//TODO

```
#### tmux
- Repo: TODO

Example configuration:
```

//TODO

```
#### tilix
- Repo: TODO

Example configuration:
```

//TODO

```
#### vim
- Repo: TODO

Example configuration:
```

//TODO

```
#### vis
- Repo: TODO

Example configuration:
```

//TODO

```
#### vscode
- Repo: TODO

Example configuration:
```

//TODO

```
#### windows-command-prompt
- Repo: TODO

Example configuration:
```

//TODO

```
#### xcode
- Repo: TODO

Example configuration:
```

//TODO

```
#### xfce4-terminal
- Repo: TODO

Example configuration:
```

//TODO

```
#### xresources
- Repo: TODO

Example configuration:
```

//TODO

```
#### xshell
- Repo: TODO

Example configuration:
```

//TODO

```
#### zathura
- Repo: TODO

Example configuration:
```

//TODO

```

## Contributing
I hacked this project together in a weekend and it grew to be bigger than
expected. The code quality could be way better and even though it is already
pretty usable, you might find bugs or other issues. The documentation is
work-in-progress.

Issues, bug-reports, pull requests or ideas for features and improvements are
**very welcome**. Also it would be great if users of specific applications can
document the usage of their respective templates, as I don't use all of them and
can't/won't test the integration for every single application.

