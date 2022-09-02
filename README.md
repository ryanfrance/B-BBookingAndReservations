# Mock B&B business - Bookings and Reservations

## Table of Contents

- [About](#about)
- [Getting Started](#getting_started)
- [Installation](#installing)
- [Notes](#notes)

## About <a name = "about"></a>

This is a mock project to learn about golang and put something into practice.

## Getting Started <a name = "getting_started"></a>

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See [deployment](#deployment) for notes on how to deploy the project on a live system.

## Installation <a name = "installing"></a>

Clone down then compile and run with: `go run ./web/cmd` <br>
This will spin up a http server viewable from localhost and whatever port you specify in `main.go` (default is `:8080`)

## Notes <a name = "notes"></a>

- Built in Go 1.19
- Uses the <a href="https://github.com/go-chi/chi/v5">Chi router</a>
- Uses <a href="https://github.com/alexedwards/scs/v2">Alex Adwards SCS session management</a>
- Uses <a href="https://">Nosurf</a>
