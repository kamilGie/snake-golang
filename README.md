# Project: Snake in Go

## Project Description

This project is an implementation of the Snake game written in Go. The application is divided into  main  file and packages
snake :

1. **main** - Responsible for rendering the game on the screen and handling keyboard input.
2. **snake** - An independent package that contains the core game logic. It serves as a facade, allowing control of the game through specific methods.

## Key Methods in the `snake` Package

- **New()** - Create environment with passed width and height
- **TakeAction()** - executes an action in the game every tick, such as moving the snake or handling collisions.
- **GetState()** - retrieves the current state of the game, including the snake's and fruit position, and game-over status.

The separation of the `snake` package is intentional, allowing it to be used independently for machine learning research in the future.
