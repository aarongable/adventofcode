# adventofcode

This is where I'll be throwing all of my
[Advent of Code](https://adventofcode.com/) work. Spoilers if you're doing
it yourself, I guess.

Because I assume that code reuse will be valuable, the repository is divided
by language at the top level, with years and days nested underneath each
language directory. This allows (for example) all of the Go code to be a single
module containing multiple packages, so that a Day 5 solution can import the
Day 1 solution as a dependency.
