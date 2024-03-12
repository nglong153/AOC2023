# Start of Day2

## Part 1

- Summary of the problem: The are maximum of 12 red, 13 green, and 14 blue. Each line is a game with an Id, return the sum of ids of all the valid game
  - What is a valid game: A game has many sets. It is valid if all the sets is valid. One set is valid if any number of the color <= the maximum number (give above)
- Solution: Read the file line by line, a game has format like this "Game n: a red, b blue, c green; ..."
  - Split the line into two part, the game ID part and the sets part
  - TrimPrefix with "Game " to get the id number
  - Split up the sets with ";", then splits up again with "," to validate the number of each color

## Part 2

- Summary of the problem: For every games, find the maximum of each color : red, green, and blue;
- Solution: Read the file line by line, a game has format like this "Game n: a red, b blue, c green; ..."
  - Split the line into two part, the game ID part and the sets part
  - TrimPrefix with "Game " to get the id number
  - Update the maximum number
