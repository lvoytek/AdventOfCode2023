# AdventOfCode2023
Advent of Code 2023 solutions written in Go

## How to use
### Create dayX from template
To create a new day folder from the template, run

```bash
make day DAY=<Day Number>
```

or

```bash
make DAY=<Day Number>
```

If no day is provided, it will default to 1.

### Run a solution
To run one or both solutions for a day, run the following

```bash
make run DAY=<Day Number> PART=<Part Number>
```

If no part number is provided, then both will be run.

### Test a day's solutions

To test the solutions for a day using the provided examples, run the following.

```bash
make test DAY=<Day Number>
```