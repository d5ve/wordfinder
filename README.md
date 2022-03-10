# WORDFINDER API v0.3

Simple perl/dancer2 JSON API to find words matching a pattern.

When supplied a list of letters, return a list of all the words that can be
built from just those letters. The input letters can include repeats. Not all
input letters need to be used. Each letter in the input list can appear at most
once in the output words.

[dockerhub: d5ve/wordfinder](https://hub.docker.com/r/d5ve/wordfinder)

## Run from dockerhub

```bash
$ docker run -d -p 8080:80 d5ve/wordfinder
```

## Run with docker

```bash
$ docker build -t wordfinder .
...
$ docker run -d -p 8080:80 wordfinder
```

## Run directly

```bash
$ cpanm Dancer2 JSON::XS
...
$ plackup --port 8080 wordfinder.psgi
```
## Test with curl
```bash
$ curl http://localhost:8080/ping
OK

$ curl http://localhost:8080/wordfinder/dgo
[ "do", "dog", "go", "god", "o" ]
```

## TODO

- Make app into a package for ease of testing.
- Tests.
- Better algorithms. Too many nested loops currently. Use benchmarking to compare different ideas.
- Caching. Normalise the input letters and have a small round-robin memory cache.
- Better error handling.

## LICENSE

Wordfinder is free software; you can redistribute it and/or modify it under the same terms as perl itself.
