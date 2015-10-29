# readinglist

## Installation
```bash
go get github.com/andrewstuart/readinglist
```

## Usage

```bash
$ readinglist ls

1.  https://news.ycombinator.com
2.  https://github.com/andrewstuart/readinglist

$ readinglist 1
# opens link number 1 (as numbered by ls)

$ readinglist pop
# opens latest link in your browser and removes it from the list

$ readinglist shift
# opens oldest link in your browser and removes it from the list

```

Don't like typing readinglist out?

```bash
ln -s $GOPATH/bin/{readinglist,rl}
```
