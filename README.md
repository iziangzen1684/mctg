# Overview
A simple Markov chain text generator written in Go
# Requirement
 + Go (>= 1.18)
 + Git
# Installation
Linux/MacOS:
```sh
git clone https://github.com/iziangzen1684/mctg
cd mctg
go install
```
Alternatively:
```sh
go install github.com/iziangzen1684/mctg
```

Add this to your shell's configuration file and reload it:
```sh
export PATH=$PATH:$(go env GOPATH)/bin
```
# Usage
The default input file is ```./data.txt```, you can change it with:
```sh
mctg -d <filename>
```
The default order length is 2, you can change it with:
```sh
mctg -O <length>
```
The default token count is 100, you can change it with:
```sh
mctg -c <token_count>
```
To display the help message:
```sh
mctg -h
```
Have fun =)
