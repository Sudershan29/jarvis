# Introduction

## Motivation

    This will be an app dedicated to help develop, manage and master skills. The name of the project was carefully researched and named after 2 seconds of thinking 
    
    
    (NOTE: The project will undergo exactly 14,000,605 name changes)

## Tech Stack

1. Backend : GoLang
2. Frontend : Hmmm

## Story

![alt text](https://media.tenor.com/E5mPvZRwZ7gAAAAd/kung-fu.gif "Kung Fu Panda | Peach Tree")

# Contribution

## Running the application

```
source scripts/init.sh
source scripts/start.sh
```

## Database

1. To add models to the schema

```
go run -mod=mod entgo.io/ent/cmd/ent new <Model>
```

2. To autogenerate models 

```
go generate ./ent
```