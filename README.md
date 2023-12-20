# Tech Test Brief

## Overview

This project gets all the running processes on your PC and updates the frontend in real time.

## Project Overview

### Tech Stack

- Golang for backend
- Vue3 and Pinia(Vuex 5) for the frontend and bundled with Vite
- Sqlite for database
- Docker compose for orchestration

1. Install dependencies

```
yarn install
```

### Starting the app

```
docker-compose up -d
```

### Project Src Structure

### Backend

```
backend/
    └── configs/                        <-- Db setup and run migration
    └── handlers/                       <-- Route handlers
    └── migrations/                     <-- Migration scripts
    └── mocks/                          <-- Test mocks
    └── mocks/                          <-- Fetches all the running processes
        ├── process_darwin.go           <-- build tags for darwin OS(e.g macos)
        ├── process_linux.go/           <-- build tags for unix OS(e.g linux , ubuntu)
        ├── process.go/                 <-- default process file
    ├── routes/                         <-- Resource routes
    ├── types                           <-- Global struct types
    ├── utils                           <-- Helper functions
    ├── .dockerignore                   <-- Docker ignore
    ├── .env.example                    <-- .env example
    ├── .gitignore                      <-- ignored files and folders
    ├── cron.go                         <-- Cron file
    ├── Dockerfile                      <-- Docker image for the backend service
    ├── go.mod
    ├── go.sum
    └── main.go

```

#### Run Tests

```
    go test ./...
```

### Frontend

```
frontend/
    └── src/
        ├── assests/                    <-- Frontend assests
        ├── components/
        ├── internals/
        ├── mocks/
        ├── router/                     <-- page router
        ├── stores/                     <-- vuex stores
        ├── styles/
        ├── utils/                      <-- Helper functions
        ├── views/                      <-- pages by modules
        ├── App.vue                     <-- Vue entry point
        ├── main.ts                     <-- Application entry point
        ├── tsconfig.json/
    ├── ***                             <-- All external files are configs for (Vitest , Tailwing , vite and tsconfig and package manager)
```

### TODOs

- Add graph showing user against total cpu and memory usage (API ready: /reports)
- Add more test coverage
- Get processes on windows PC
