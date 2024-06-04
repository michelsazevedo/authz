# Authz

Authz is a microservice for managing user authentication. It supports user registration, login, and token-based authentication using JSON Web Tokens ([JWT](https://jwt.io/)).

## Built With

- [Go](https://golang.org/)

Plus *some* of packages, a complete list of which is at [/master/go.mod](https://github.com/michelsazevedo/authz/blob/master/go.mod).

## Instructions

### Running with Docker
[Docker](www.docker.com) is an open platform for developers and sysadmins to build, ship, and run distributed applications, whether on laptops, data center VMs, or the cloud.

If you haven't used Docker before, it would be good idea to read this article first: Install [Docker Engine](https://docs.docker.com/engine/installation/)

1. Install [Docker](https://www.docker.com/what-docker) and then [Docker Compose](https://docs.docker.com/compose/):

2. Run `docker compose build --no-cache` to build the images for the project.

3. Finally, run the local app with `docker compose up web` and authz will perform requests.

4. Aaaaand, you can run the automated tests suite running a `docker compose run --rm test` with no other parameters!

## License
Copyright © 2024
