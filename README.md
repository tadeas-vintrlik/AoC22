# Advent of Code 2022

[Advent of Code](https://adventofcode.com/) is an annual set of programming puzzles.
Some might call it competitive programming. I mostly use it to learn new languages and improve
my knowledge of the ones I already know.

This is my third year doing it. In 2019 I discovered it, did one day and the forgot about it. Every
year since 2020 I've given it my best. 2020 was C. 2021 was learning Go. I wrote a lot of Go recently
and looking back it was not really the best Go. This year I decied to do Go again but this year I
hope for it to be idiomatic Go.

## Project structure

- the project follows a very typical golang project structure

| directory | explanation                                                                 |
| --------- | --------------------------------------------------------------------------- |
| cmd       | directory containing executables                                            |
| internal  | directory of project-only libraries                                         |
| pkg       | directory of libraries used by this project but that can imported elsewhere |

# Solutions

- an executable providing the solution for all days can be found in `cmd/solutions`

```sh
cd cmd/solutions
go run . # To run a temporary executable
go build # To create an executable instead
```

## Testing

- the puzzle text makes it easy to create unit tests
- to run all of them:

```sh
go test ./...
```
