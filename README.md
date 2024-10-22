# Process Monitor

![](https://github.com/dagenius007/golang-process-logs/blob/main/ProcessLogs.gif)

## Overview

This project gets all the running processes on your PC and updates the frontend in real time.

## Project Overview

### Tech Stack

- Golang for backend
- Vue3 and Pinia(Vuex 5) for the frontend and bundled with Vite
- Sqlite for database
- Docker compose for orchestration

### Starting the app

Build the project

```
docker-compose build

```

Start the project

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

- Add more test coverage
- Get processes on windows PC
- Get full command as -comm limits to 16 characters

### References

- https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/6/html/deployment_guide/ch-proc#:~:text=The%20%2Fproc%2F%20directory%20(also,kernel's%20view%20of%20the%20system.
- https://man7.org/linux/man-pages/man5/proc.5.html
- https://sites.ualberta.ca/dept/chemeng/AIX-43/share/man/info/C/a_doc_lib/cmds/aixcmds4/ps.htm#A163C1169
- https://man7.org/linux/man-pages/man1/ps.1.html
-
