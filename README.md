# Golf IndX

Golf Handicap Tracker

by Tim Robillard
hello@timrobillard.ca

## Installation

Clone the repo, create a `.env` file from `.env.example`, and run the code.

```bash
cp .env.example .env
make run
```

## Local Development

For hot reloading, it takes 3 shells running at once, which can be started with the following commands

```bash
# for updating public/styles.css
make tailwind
```

````bash
# for generating the templ.go files
make templ
```bash
# for running the server
air
````

_NOTE_ the proxy running in `make templ` is unstable, and may crash occasionally. Just restart it and air and it should work again!
