# Magical Arena

## Problem Statement
Design a Magical Arena. Every Player is defined by a “health” attribute, “strength” attribute and an “attack” attribute - all positive integers. The player dies if his health attribute touches 0. 

After Reading the Problem Statement throughly I decided to do this project using Go

## Overview
I came up with this folder structure in order to complete the project
```
Folder Structure:
Magical Arena/
├── cmd/
│   └── magicalarena/
│       └── main.go
├── internal/
│   ├── player/
│   │   ├── player.go
│   │   └── player_test.go
│   ├── match/
│   │   ├── match.go
│   │   └── match_test.go
│   ├── utils/
│       ├── constants/
│       │   └── constants.go
│       ├── dice/
│       │   └── dice.go
│       └──  helpers/
│           └── helpers.go
├── go.mod
├── Problem Statement.txt
└── README.md
```

## Getting Started
To run the Magical Arena simulation, follow these steps:

1. Clone or download the project repository to your local machine.
2. Open the project in your preferred IDE such as VS Code or Goland.


## Running the Application
1. Move to the directory Magical Arena
```bash
cd ~/pathtorepository/Magical Arena
```
2. Command to run main.go
```bash
go run ./cmd/magicalarena/main.go
```
3. Command to run test files
```bash
go test ./...
```