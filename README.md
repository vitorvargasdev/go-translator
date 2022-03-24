# About
A simple Go project made to get translations from Google Translate.

This project was developed mainly using google technology called chromedp which is used to browse google chrome using programming.

# Requirements
- Golang
- Google Chrome

# Installation Guide

To download, compile, and install the latest release, do this:
```sh
git clone https://github.com/instalando/go-translator.git
cd go-translator
go install
go build
```

# Usage

After the build you will have the executable in the same folder called "gotranslator"

```sh
$ ./gotranslator

NAME:
Go Translator ~ Developed by @instalando - Translate any text to any language

USAGE:
main [global options] command [command options][arguments...]

COMMANDS:
translate Translate text
help, h Shows a list of commands or help for one command

GLOBAL OPTIONS:
--help, -h show help
```

To use it, just add the parameters inside the translate function that are "to" and "from" defining the language you are using and the language you want to translate. by default "from" is "auto".

```sh
$ ./gotranslator translate --text="Hello World!" --from="en" --to="ja"
```

And the result should be: "こんにちは世界"