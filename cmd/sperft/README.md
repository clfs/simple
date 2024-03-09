# sperft
The `sperft` tool calculates Perft counts in different positions.

It ignores draws by:

- repetition
- the fifty-move rule
- the seventy-five-move rule
- insufficient material

TODO: Write `sperft` tests.

## Install

```text
go install github.com/clfs/simple/cmd/sperft@latest
```

## Uninstall

```text
rm -i $(which sperft)
```

## Usage

```text
$ sperft -h
Usage of sperft:
  -depth int
        search depth (default 1)
  -fen string
        position (default "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
```

## Example

In position `2r2k2/1pnR2p1/p7/P1PP4/5pp1/8/5PPP/6K1 w - - 0 39`, there are 4,009
possible leaf nodes after 3 plies, and 58 of these leaf nodes started with White
playing `d7d8`.

```text
$ sperft -depth 3 -fen "2r2k2/1pnR2p1/p7/P1PP4/5pp1/8/5PPP/6K1 w - - 0 39"
4009
c5c6    319
d5d6    274
d7c7    258
d7d6    360
d7d8    58
d7e7    395
d7f7    46
d7g7    326
f2f3    314
g1f1    332
g1h1    298
g2g3    313
h2h3    366
h2h4    350
```
