# Coldtrail (Go)

A number-guessing duel for the terminal, written in Go (standard library
only, no external modules).

A number from 1 to 100 is hidden. You get 7 guesses per round, and after
each guess the game reports a temperature — freezing, cold, cool, warm,
hot, burning hot — plus whether the trail leads higher or lower. Three
rounds; leftover guesses convert to points.

## Build

```sh
go build -o coldtrail .
```

## Play

```sh
./coldtrail             # random game
./coldtrail --seed 42   # reproducible game (fixed hidden numbers)
```

Ratings run from Lost up to Trailblazer (a perfect game scores 420).

## Tested envelope

- Built clean with Go 1.23 (`go build`, `go vet`).
- Verified end to end with scripted stdin playthroughs: a perfect
  3-round game (score 420, rating Trailblazer) and an out-of-guesses
  path, both reaching the game-over screen, using `--seed 42`.
