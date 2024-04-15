# advent-of-code

My solutions to the Advents of Code. Probably the best programming challenge can be find [here](https://adventofcode.com/).

![Progress2021](https://progress-bar.dev/10/?scale=25&title=2021&width=120&suffix=/25)

![Progress2022](https://progress-bar.dev/11/?scale=25&title=2022&width=120&suffix=/25)

![Progress2023](https://progress-bar.dev/4/?scale=25&title=2023&width=120&suffix=/25)

## .env

Task has a functionality to download puzzle input. To be able to do that one must login to advent of code page and extract his cookie session id. Then set it on `.env` file like this:

```bash
ADVENT_OF_CODE_COOKIE=53616...foobar...8fea6cc
```

## new solution

New solution should start with command like `task new -- 2021 1`. First argument is `year` and second `day` of challenge. This will create new folder in appropriate folder. Extract template into and download puzzle input.

## prove solution

One solution is made execute `task test`. All tests should pass.
