# Table of Contents
1. [What is this?](#what-am-i-looking-at)
2. [How to use this](#how-to-use-this)
3. [Examples of usage](#examples-of-usage)
4. [How to compile it](how-do-i-compile-this)

## What am I looking at
It's a CLI tool that I made for finding the absolute path for a subdirectory from a certain path so that I can CD to it easily. If you have golang installed, you can build the directory then place its path into your bash/zsh/etc. profile under the path directory so you can call it everywhere.

## How to use this
Just call it on a folder name. There are three possible flags:
* -l - int; limit how many results are retrieved. Default: -1 (infinite)
* -w - string; set the base directory. Default: current working directory
* -r - bool; retrieve relative paths instead of absolute paths. Default: false

NOTE: If any flags are set, the target folder name come after all flags

## Examples of usage
```
way build
way -r -l 5 -w /User/me node_modules
```

## How do I compile this
* Install the Golang runtime (https://go.dev)
* Navigate to the folder with the project in it
* Run: go build way.go

## How do I add this to my path variables?
* Open your terminal profile file (mine is .zshrc)
* Add the path variable on a blank line, like:

```
export PATH="$PATH/route/to/compiled/binary"
```

* OR you can just add it to your usr/bin folder if you're on mac/unix and don't already have a way binary (and you like to live dangerously)

## Planned Updates
* At some point I may add my own logic for accessing subdirectories using coroutines instead of the built-in fs.WalkDir

## Changelog
* 1/15/2022: First release