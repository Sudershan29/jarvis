# Introduction

## Aim

1. This is an app that is aimed to improve productivity by helping you scheduling your everyday activities like tasks, meetings, etc.
2. The app offers intergration with your google calendar, and automatically propose a best schedule to complete your tasks.

## Motivation

1. Improve productivity, and set realistic goals for the year
2. Plan events by looking at multiple people's calendars

## Tech Stack

1. Backend  : GoLang
2. Frontend : ReactJs

## Story

![alt text](story.gif "Kung Fu Panda | Peach Tree")

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

## Testing

```
go test ./tests/...
```