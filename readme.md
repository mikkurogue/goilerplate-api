# Goilerplate - Powered by echo

The boilerplate for your Go API!

## Why?

Why not, its just a fun side project for me on the side. It's ease of access too for anyone wanting to get started with a boilerplate for an api.

## Lightweight-ish?

I want to try and keep this boilerplate to as minimal amount of direct deps as possible, for transparency the reasoning is because then no one should feel pidgeon-holed into using some potentially obscure dependencies that they potentially may not need.

In terms of cli styling for the binary, I would be open to using things like charmbracelet for logger purposes to make things look _cool_ but at this current point in time I don't see any value as of yet to do so.

## What comes out of the box?

- Relatively simple project structure
- 2 GET requests that send a string and a json object
- 2 POST requests - multipart and json body
- Basic JWT token generation
  - With example /login post route
- Example protected routes for auth
- Mongodb driver
- Support for turso


### Database caveats
I do highly recommend uncommenting/removing the database files you do not need for your use case. For instance, if you are not planning on using Turso, then make sure to remove the Turso boilerplate. The same goes for mongo db.

The reason I have them both in the main branch is because I want to keep everyones options open and allow people to work with whatever they want. The point of this boilerplate is to make it as robust, lightweight and adaptable as possible.

## How to run?

Clone the repository and simply navigate to it with your terminal and fire the command `go run main.go`
To build the binaries run `go build .`

Then make a call to `localhost:1337/hello-world` to get an initial GET response!

To test the POST requests, I recommend using a client like Postman or Insomnia.
