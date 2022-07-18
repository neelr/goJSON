# goJSON

[![Run on Repl.it](https://repl.it/badge/github/neelr/gojson)](https://repl.it/github/neelr/gojson)

<img src="https://gojson.hacker22.repl.co/gojson.png" width="100"/>

A simple and lightning fast database written in Golang!

This api is meant for small projects that just need a fast database, and dont want to mess around with tokens and setting one up yourself! Just copy the link, and use the API methods to store and retrieve any JSON you need!

Open Source and accepting contributors!

## API

**UPDATE**: To check the magnitude of requests, I made an open source log for the API. Check at https://db.neelr.dev/logs. It only tracks the number of requests, not IP's or anything else.

Use the UI to take a look at your token's database: https://dash.db.neelr.dev

Your very own database API endpoint is `/api/:key`

### Add/Update Data

**POST** or **PUT** `/api/f0e4408dc81b90365ed5b11112ff2575`!

Also you can update inner objects using keys!

ex. POST `/api/f0e4408dc81b90365ed5b11112ff2575/hello/hi`

```json
{
	"go": true
}
```

will update it to create a JSON

```json
{
	"hello": {
		"hi": {
			"go": true
		}
	}
}
```

**IMPORTANT UPDATE:**  
Take note that **PUT** method only updates, and does not delete any keys in the database when writing, while **POST** overwrites any existing keys in the same directory. For the example above both work, but instead, if there were other values at /hello/hi **POST** would overwite all of them with `go:true`

### View Data

**GET** `/api/f0e4408dc81b90365ed5b11112ff2575` or navigate through the keys with `/api/f0e4408dc81b90365ed5b11112ff2575/hello` would return

```json
{
	"hi": {
		"go": true
	}
}
```

### Delete Data

**DELETE** `/api/f0e4408dc81b90365ed5b11112ff2575` and you can delete using the same key navigation given above!

## Building and Running

1. Install Golang on your computer

2. Run `git clone https://github.com/neelr/gojson`

3. Create an empty `database` folder

4. Go into the directory and run `go run .`!

## Contribution

Any type of feedback, pull request or issue is welcome.

**MIT LICENSE**

