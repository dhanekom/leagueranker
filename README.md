# Leagueranker
## Overview
Takes league game scores as a text input and outputs teams ordered by their rank. Input data can be piped into the application or a input file can be used.
## Examples
### Pipe data into the application
```console
$ cat input.txt | ./leagueranker
```
or
```console
$ ./leagueranker "Lions 3, Sharks 2"
```
### Use an input file
```console
$ ./leagueranker -file input.txt
```