# gator

A barebone blog aggregator in Go

Made in the context of the [Boot.dev class: Build a Blog Aggregator in Go](https://www.boot.dev/courses/build-blog-aggregator-golang)

## How to use

1. Make sure you have [Postgresql](https://www.postgresql.org/) and [Go](https://go.dev/) installed.

2. Install `gator`by executing `go install github.com/Lauloque/gator@latest`

3. Make sure you have a running postgresql server, then setup a config file, .gatorconfig.json in your home directory. That file should contain the url to your postgresql server. Example:

```json
{ "db_url": "postgres://username:@localhost:5432/database?sslmode=disable" }
```

## Commands

Global usage: `gator <command> [arguments...]`

- `register [name]` - Register a new user

- `login [name]` - Log in as an existing user

- `users` - List all users

- `addfeed [feed name] [feed url]` - Adds a new feed to the database

- `feeds` - Displays feeds that have been added to gator

- `agg [duration]` - Aggregates added feeds at the given interval duration eg: '1s', '1m', '1h' etc

- `follow [feed url]` - Follows a feed already in the database

- `follows` - Display feeds that the current user is following

- `unfollow [feed url]` - Stops following a feed

- `browse [number]` - Displays `[number]` most recent posts from followed feeds

- `reset` - Deletes all users from database
