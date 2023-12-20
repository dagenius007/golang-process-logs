# Tech Test Brief

## Overview

We are seeking a capable developer to build a front-end application using either React or Next.js, combined with TypeScript. Your choice of framework should be based on your preference and expertise. The application should closely align with the designs found in this [Figma file](https://www.figma.com/file/hGBI3zpyHia5yrWsgeMP3K/Untitled?node-id=0%3A1&mode=dev). It is essential that the application utilises Styled Components for styling.

## Project Overview

### Tech Stack

React was used for building the application. Follow the steps to setup on the application

1. Install dependencies

```
yarn install
```

### Starting the app

Now you are ready to run the application.:

#### Starting the app with local apis:

```
yarn start
```

### Project Src Structure

#### Components

This contains all styled components. It is setup in a block format where all the dependency are inside the folder. Dependencies like the styles , types and hooks(if need be). The project is quite small and so all the dependencies are all in the same file.

#### Global Styles

All generic styles used in various parts of the project for basic styling. It is also have media query for responsive design.

#### Hooks

Generic functions used by various components to improve performance and also for abstractions of function

#### Internals

General configurations used in the project

#### Mocks

Mocked data for tests.

#### Pages

Representation of each modules. It will be used to determine component loaded on a particular route

#### Types

Generic interface to define the structure and shape of an object.

#### Utils

They are helper functions.
