# Link Shortener (Go)

This is a simple URL shortener written in Go. I put it together as a way to get to grips with the basics of the language. If you’re expecting a full-blown production service, you may want to lower your expectations. This is more of a learning exercise.

## What does it do?

You send it a long URL, it gives you a short code. If the original link doesn’t work, the code quietly disappears.

## How does it work?

- POST to `/create` with a JSON body like `{ "url": "https://example.com" }`.
- You’ll get back a short code.
- Visit `/r/{short_code}` to be redirected.
- If the original URL is broken, the code is removed. No fuss.

## Why Go?

I wanted to see what all the fuss was about. Go is supposed to be fast, simple, and good for web. This project is my way of poking it with a stick to see what happens.

## Anything else?

There’s not much in the way of error handling, security, or persistence. If you restart the server, all your links are gone. If you want something more robust, you’ll need to add a database and a bit more patience.

If you spot something odd, it’s probably on purpose. Or it’s a bug. Either way, fun stuff.
