<h1 align="center">
   <a href="#"> Code challenge Levee API </a>
</h1>

<h3 align="center">
    Backend Engineer Code Challenge - Levee
</h3>

<p align="center">

<img alt="Testing" src="https://github.com/fabianoleittes/code-challenge-levee/workflows/Tests%20&%20Linters/badge.svg?branch=main">

  <img alt="GitHub language count" src="https://img.shields.io/github/languages/count/fabianoleittes/code-challenge-levee?color=%2304D361">

  <img alt="Repository size" src="https://img.shields.io/github/repo-size/fabianoleittes/code-challenge-levee">

  <a href="https://github.com/fabianoleittes/code-challenge-levee/commits/main">
    <img alt="GitHub last commit" src="https://img.shields.io/github/last-commit/fabianoleittes/code-challenge-levee">
  </a>

   <img alt="License" src="https://img.shields.io/badge/license-MIT-brightgreen">
   <a href="https://github.com/fabianoleittes/code-challenge-levee/stargazers">
    <img alt="Stargazers" src="https://img.shields.io/github/stars/fabianoleittes/code-challenge-levee?style=social">
  </a>
</p>


<h4 align="center">
	 Status: WIP
</h4>

<p align="center">
 <a href="#about">About</a> •
 <a href="#features">Features</a> •
 <a href="#layout">Layout</a> •
 <a href="#how-it-works">How it works</a> •
 <a href="#tech-stack">Tech Stack</a> •
 <a href="#contributors">Contributors</a> •
 <a href="#author">Author</a> •
 <a href="#user-content-license">License</a>

</p>

## About

This project is a simple API for some `Job` routines, such as creating, listing, and activate.

---

## Features

- [x] Create job
- [ ] List the all jobs
- [ ] Activate the status for a specific job
- [ ] List the percentage and number of active jobs by category.

---

## How it works

This project is Restful API:
1. Backend


### Pre-requisites

Before you begin, you will need to have the following tools installed on your machine:
[docker](https://docs.docker.com/install/), [docker compose](https://docs.docker.com/compose/install/), [Git](https://git-scm.com).

In addition, it is good to have an editor to work with the code like [VSCode] (https://code.visualstudio.com/)

#### Running the Backend (server)

```bash

# Clone this repository
$ git clone https://github.com/fabianoleittes/code-challenge-levee

# Access the project folder cmd/terminal
$ cd code-challenge-levee

# Environment variables
$ make init

# Run the application in development mode
$ make up

# The server will start at port: 3001 - go to http://localhost:3001

# View logs
$ make logs
```


#### API Request

 Endpoint        | HTTP Method           | Description       |
| --------------- | :---------------------: | :-----------------: |
| `/v1/jobs` | `POST`                | `Create jobs` |
| `/v1/health`| `GET`                 | `Health check`  |
---

## Tech Stack

The following tools were used in the construction of the project:

#### **API**  ([Golang](https://golang.org/))

-   **[Gorilla/mux](https://github.com/gorilla/mux)**
-   **[Gin Web Framework](https://github.com/gin-gonic/gin)**
-   **[PostgreSQL](https://www.postgresql.org/)**
-   **[MongoDB](https://www.mongodb.com/)**

**Utilities**


-   Commit Conventional:  **[Commitlint](https://github.com/conventional-changelog/commitlint)**
-   API Test:  **[Insomnia](https://insomnia.rest/)**
---


## How to contribute

1. Fork the project.
2. Create a new branch with your changes: `git checkout -b my-feature`
3. Save your changes and create a commit message telling you what you did: `git commit -m" feature: My new feature "`
4. Submit your changes: `git push origin my-feature`
> If you have any questions check this [guide on how to contribute](./CONTRIBUTING.md)

---

## Author

<a href="https://fabianoleittes.me/">
 <img style="border-radius: 50%;" src="https://avatars3.githubusercontent.com/u/279344?v=4" width="100px;" alt=""/>
 <br />
 <sub><b>Fabiano Leite</b></sub></a>
 <br />

[![Twitter Badge](https://img.shields.io/badge/-@fabianoleittes-1ca0f1?style=flat-square&labelColor=1ca0f1&logo=twitter&logoColor=white&link=https://twitter.com/fabianoleittes)](https://twitter.com/fabianoleittes) [![Linkedin Badge](https://img.shields.io/badge/-Fabiano-blue?style=flat-square&logo=Linkedin&logoColor=white&link=https://www.linkedin.com/in/fabianoleittes/)](https://www.linkedin.com/in/fabianoleittes/)

---

## 📝 License

This project is under the license [MIT](./LICENSE).

##### Made with love by Fabiano Leite 👋🏽 [Get in Touch!](Https://www.linkedin.com/in/fabianoleittes/)
---
