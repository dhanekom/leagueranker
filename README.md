# Leagueranker
## Overview
Takes league game scores as a text input and outputs teams ordered by their rank. Input data can be piped into the application or a input file can be used.

Sample input:
```console
Lions 3, Snakes 3
Tarantulas 1, FC Awesome 0
Lions 1, FC Awesome 1
Tarantulas 3, Snakes 1
Lions 4, Grouches 0
```
Sample output:
```console
1. Tarantulas, 6 pts
2. Lions, 5 pts
3. FC Awesome, 1 pt
3. Snakes, 1 pt
5. Grouches, 0 pt
```
## Build instructions
The steps below assume that you are running a Linux or Mac and that you already have git and and the go language installed.

Step 1: Clone the repo
```console
Create a directory for your project (e.g. my-project)
cd into the directory (e.g. cd my-project)
git cone https://github.com/dhanekom/leagueranker.git
```
Step 2: Build the application
For the root directory op your project (e.g. the my-project directory)
```console
cd cmd/leagranker
go build
```
## Usage examples
### Pipe data into the application
from the my-project/cmd/leagranker directory
```console
$ cat input.txt | ./leagueranker
```
### Using an input file
```console
$ ./leagueranker -file input.txt
```