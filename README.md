# advent-of-code

My solutions to the Advents of Code. Probably the best programming challenge can be find on [adventofcode.com](https://adventofcode.com/).

![Progress2021](https://progress-bar.xyz/11/?scale=25&title=2021%2011/25&width=120)

![Progress2022](https://progress-bar.xyz/11/?scale=25&title=2022%2011/25&width=120)

![Progress2023](https://progress-bar.xyz/4/?scale=25&title=2023%2004/25&width=120)

![Progress2024](https://progress-bar.xyz/11/?scale=25&title=2024%2011/25&width=120)

![Progress2025](https://progress-bar.xyz/1/?scale=25&title=2025%201/25&width=120)

## .env

Task has a functionality to download puzzle input. To be able to do that one must login to advent of code page and extract his cookie session id. Then set it on `.env` file like this:

```bash
ADVENT_OF_CODE_COOKIE=53616...foobar...8fea6cc
```

## new solution

New solution should start with command like `task new -- 2021 1`. First argument is `year` and second `day` of challenge. This will create new folder in appropriate folder. Extract template into and download puzzle input.

## prove solution

One solution is made execute `task test`. All tests should pass.
