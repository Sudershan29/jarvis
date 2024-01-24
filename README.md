# Introduction

## Aim

1. This app is designed to enhance productivity by assisting you in scheduling your daily activities such as tasks and meetings.
2. The app integrates seamlessly with your Google Calendar and automatically suggests the optimal schedule for completing your tasks.

## Motivation

1. Boost productivity and set achievable goals for the year.
2. Coordinate events by checking multiple people's calendars.

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