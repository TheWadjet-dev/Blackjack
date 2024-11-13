# Blackjack Game in Go

This is a simple web-based Blackjack game built with Go. The game logic is handled on the server side in Go, and the interface is served through HTML. This project demonstrates basic web server functionality and game mechanics for a classic Blackjack game.

## Project Structure

- **main.go**: Contains the main game logic and the HTTP server code.
- **index.html**: The frontend HTML file, which interacts with the server to display the game and handle player actions.
- **Dockerfile**: Configuration to containerize the application with Docker.

## Prerequisites

To run this project, you need:
- [Go](https://golang.org/dl/) (version 1.20 or later)
- [Docker](https://www.docker.com/get-started) (optional, if you want to run the application in a container)

## Running the Application Locally

1. **Clone the repository**:
    ```bash
    git clone https://github.com/TheWadjet-dev/Blackjack.git
    cd blackjack-go
    ```

2. **Run the Go application**:
    ```bash
    go run main.go
    ```

3. **Access the game**:
   Open your browser and go to `http://localhost:8080` to start playing the game.

## Running the Application with Docker

1. **Build the Docker image**:
    ```bash
    docker build -t blackjack-go .
    ```

2. **Run the Docker container**:
    ```bash
    docker run -p 8080:8080 blackjack-go
    ```

3. **Access the game**:
   Open your browser and go to `http://localhost:8080`.

## How to Play

- The game is a classic version of Blackjack. 
- Upon starting a new game, you are dealt two cards, and the dealer is dealt two cards (one face down).
- You can choose to **Hit** (draw a card) or **Stand** (end your turn).
- The goal is to get as close to 21 points as possible without going over.
- If your score exceeds 21, you "bust" and lose the game.
- Once you stand, the dealer will reveal their hand and draw cards until they reach 17 or higher.

## Game Logic Overview

- **Deck and Card Generation**: The game creates and shuffles a standard 52-card deck.
- **Scoring**: Aces count as 1 or 11 points, depending on the current hand score.
- **Dealer Logic**: The dealer must hit until reaching a score of at least 17.
