# Magical Arena

## Problem Statement
Design a Magical Arena. Every Player is defined by a “health” attribute, “strength” attribute and an “attack” attribute - all positive integers. The player dies if his health attribute touches 0. 

After Reading the Problem Statement throughly I decided to do this project using Go

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
│   │   ├── helpers/
│   │   │   └── helper.go
│   │   ├── constants/
│   │   │   └── constants.go
│   │   └── dice/
│   │       └── dice.go
│   └── mod.go
├── go.mod
├── Problem Statement.txt
└── README.md
```
