# Connect 4 AI in GoLang

## Overview

This repository contains an implementation of the classic game Connect 4 in the Go programming language. The game allows two players to take turns dropping colored coins into a vertically suspended grid. The first player to connect four of their coins in a row, either horizontally, vertically, or diagonally, wins the game.

## Features

- **Human vs Human Gameplay:** Play against a friend by taking turns to make moves on the game board.
- **Human vs AI Gameplay:** Challenge yourself against a computer opponent that uses the minimax algorithm to make strategic moves.


## Getting Started

1. Ensure you have Go installed on your machine.
2. Clone this repository to your local machine.
3. Open a terminal window and navigate to the project directory.
4. Run the following command to execute the Connect 4 game:
   ```bash
   go run connect4.go
   ```
5. Follow the on-screen instructions to make moves and enjoy the game!

## Gameplay

- The game is played on an 8x8 grid.
- Players take turns dropping their colored discs into a column of their choice.
- The first player to connect four discs of their color in a row wins.
- The game ends in a tie if the board is full and no player has won.

## AI Implementation

The AI opponent uses the minimax algorithm with alpha-beta pruning to make intelligent moves. The algorithm explores possible future moves to find the optimal strategy while minimizing the number of nodes evaluated.

## Contributing

Feel free to contribute to the development of this Connect 4 AI in GoLang. If you have ideas for improvements, bug fixes, or additional features, submit a pull request and join the community effort.

## License

This project is licensed under the [MIT License](LICENSE). Feel free to use, modify, and distribute the code as needed.

Enjoy playing Connect 4 in GoLang!
