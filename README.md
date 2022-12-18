# My stats

*Every year but 2022 is >24h*

## 2022

|Day|Name|Time|Rank|Time|Rank|
|-|-|-|-|-|-|
|&nbsp;&nbsp;1|[Calorie counting](https://adventofcode.com/2022/day/1)|04:27:47 💤|35634|04:34:32 💤|33772|
|&nbsp;&nbsp;2|[Rock paper scissors](https://adventofcode.com/2022/day/2)|03:02:40 💤|32160|03:48:02 💤|34395|
|&nbsp;&nbsp;3|[Rucksack reorganization](https://adventofcode.com/2022/day/3)|00:38:19|11855|00:44:43|9742|
|&nbsp;&nbsp;4|[Camp cleanup](https://adventofcode.com/2022/day/4)|02:15:40 💤|21907|02:40:15 💤|22753|
|&nbsp;&nbsp;5|[Supply stacks](https://adventofcode.com/2022/day/5)|00:47:19|9299|00:51:58|8656|
|&nbsp;&nbsp;6|[Tuning trouble](https://adventofcode.com/2022/day/6)|00:10:23|6419|00:12:15|5931|
|&nbsp;&nbsp;7|[No space left on device](https://adventofcode.com/2022/day/7)|01:46:52|10534|01:51:12|9502|
|&nbsp;&nbsp;8|[Treetop tree house](https://adventofcode.com/2022/day/8)|00:24:24|4216|00:32:42|2547|
|&nbsp;&nbsp;9|[Rope bridge](https://adventofcode.com/2022/day/9)|01:01:14|8492|01:19:43|5912|
|&nbsp;&nbsp;10|[Cathode-Ray Tube](https://adventofcode.com/2022/day/10)|00:42:23|8177|01:05:06|6867|
|&nbsp;&nbsp;11|[Monkey in the Middle](https://adventofcode.com/2022/day/11)|01:11:38|6885|09:05:34|20325|
|&nbsp;&nbsp;12|[Hill Climbing Algorithm](https://adventofcode.com/2022/day/12)|08:57:59 💤|19279|09:17:25 💤|18728|
|&nbsp;&nbsp;13|[Distress Signal](https://adventofcode.com/2022/day/13)|02:06:10|7236|02:39:21|7471|
|&nbsp;&nbsp;14|[Regolith Reservoir](https://adventofcode.com/2022/day/14)|01:37:58|6631|01:51:51|6194|
|&nbsp;&nbsp;15|[Beacon Exclusion Zone](https://adventofcode.com/2022/day/15)|03:06:18|8765|05:51:24|8001|
|&nbsp;&nbsp;16|[Proboscidea Volcanium](https://adventofcode.com/2022/day/16)|11:34:21|8634|15:20:15|6299|
|&nbsp;&nbsp;17|[Pyroclastic Flow](https://adventofcode.com/2022/day/17)|05:38:50|5909|13:47:27|6531|
|&nbsp;&nbsp;18|[Boiling Boulders](https://adventofcode.com/2022/day/18)|00:33:40|3363|02:10:22|3423|

💤 = I didn't wake up :) 

# Table of contents:
1. [ My stats ](#my-stats)
2. [ What is advent of code ](#what-is-advent-of-code)
3. [ Conventions ](#conventions)
    - [ Why I don't include input files ](#why-i-dont-include-input-files)
    - [ File naming convention ](#file-naming-convention)
    - [ Commit convention ](#commit-convention)
    - [ Pull request convention ](#pull-request-convention)

# What is advent of code
[Advent of code](https://adventofcode.com/) challenges, resolved using golang

> Advent of Code is an Advent calendar of small programming puzzles for a variety of skill sets and skill levels that can be solved in any programming language you like. People use them as interview prep, company training, university coursework, practice problems, a speed contest, or to challenge each other.

# Conventions

## Why I don't include input files:
AoC creator asked to not share input files (in [this tweet](https://mobile.twitter.com/ericwastl/status/1465805354214830081) and in the [subreddit wiki](https://www.reddit.com/r/adventofcode/wiki/faqs/copyright/puzzle_texts/)), so I do not include them into my repo.

You can still run every challenge with your input, available at [Advent of code](https://adventofcode.com/), after login (input is unique for every user).

## File naming convention:
`challenge.go`: clean solution with comments

`challenge_original.go`: first solution writter, the one that calculated the result to submit

`challege_bruteforce.go`: inefficient algorithms written to test efficiency

`easyInput.txt`: input used in examples

`input.txt`: input to compute the result on

`input2.txt`: another input to test the result (usually from a smurf account or from a friend) (inputs are not the same for every user)

## Commit convention:
`YEAR-DAY: working`: usually first solution for the YEAR - DAY challenge

`YEAR-DAY: refactor`: refactor of the clean solution or refactor of the comments

`YEAR-DAY: add input`: add a test input for the YEAR - DAY challenge

`repo: ...`: adjustments to the repository (not challenge related)

## Pull request convention:
I usually create a new branch (`YEAR-DAY`) for every challenge (so for every day) and I merge (sending a PR) when it works or whenever I think it is time to merge
