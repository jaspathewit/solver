# solver
This project was created to solve Question20 of the 2022 Belgian Defensie puzzles.
[https://beldefnews.mil.be/cyber/2021_opgave_NL.pdf](https://beldefnews.mil.be/cyber/2021_opgave_NL.pdf)

Question20 was about rolling a die around on a board. Restrictions were placed on which position on the board
the dice could move to and a route had to be found such that the dice visited each position on the board and
ended up at a location a "knights" move away from the starting position.

The application works in the following way; Workers (solvers) are created and listen to a puzzle channel.
A worker takes a puzzle from the channel and calculates all valid moves of the dice.
For each valid move the board is evaluated to see if the puzzle has been solved. If it has been solved
the Solution is put onto a result channel. If it has not been solved, boards representing having
moved the dice to the valid positions are created and placed on the puzzle channel.

The application was the first time implementing a worker pool. It was also the first time that I used the 
newly introduced to Golang generics. The application also had to implement a Depth First Search to detect
and reject partitioned solutions.

The idea of splitting off the solving of each step of a problem into a worker pool
is a useful approach to solving problems like this. The repository contains "sudoku" which is an attempt to
use the same approach to solve a sudoku puzzle (Particularly the Samurai Evil level from the Saturday Times).
This also takes into account the possibility of sudokus with different topologies.

## Build and test
In the unlikely event that this code will interest anyone.

It should be sufficient to clone the repository and "go build" the program. either question20 or sudoku.

## Contribute
I would be amazed if anyone was interested in contributing... but if anyone were to feel so inclined
I'm open to pull requests.

Maybe I'll go back sometime and add some tests, and over polish it a bit.



