# readinglist
### A dead-simple CLI link management tool

## Installation
```bash
go get github.com/andrewstuart/readinglist
```

## Usage

```bash
$ readinglist push https://news.ycombinator.com
$ readinglist add https://github.com/andrewstuart/readinglist
# Add a few links to the repository

$ readinglist ls
1.  https://news.ycombinator.com
2.  https://github.com/andrewstuart/readinglist

$ readinglist 1
# opens link number 1 (as numbered by ls)

$ readinglist pop
# opens latest link in your browser and removes it from the list

$ readinglist shift
# opens oldest link in your browser and removes it from the list

$ readinglist rm 3
# remove the third link

$ readinglist splice 2
# open and remove the second link

$ readinglist git init
# Initialize a git repo to track changes

$ readinglist git ...
# Run a git command inside the local data store

$ readinglist -list linux add https://lwn.net
# Add to a specific, separate list

$ readinglist -list linux
1.  https://lwn.net

$ readinglist 1..3
#Open all links 1-3

```

Don't like typing readinglist out?

```bash
ln -s $GOPATH/bin/{readinglist,rl}
```
