# WORDFINDER API

Simple perl/dancer2 JSON API to find words matching a pattern.

When supplied a list of letters, return a list of all the words that can be
built from just those letters. The input letters can include repeats. Not all
input letters need to be used. Each letter in the input list can appear at most
once in the output words.

## Run via docker

```bash
$ docker build -t wordfinder .
...
$ docker run -d -p 8080:80 wordfinder
...
$ curl http://localhost:8080/ping
OK

$ curl http://localhost:8080/wordfinder/dgo
[ "do", "dog", "go", "god" ]
```

## Run directly

```bash
$ cpanm Dancer2 JSON::XS
...
$ plackup --port 8080 wordfinder.psgi
...
$ curl http://localhost:8080/ping
OK

$ curl http://localhost:8080/wordfinder/dgo
[ "do", "dog", "go", "god" ]
...
```

## TODO

- Tests.
- Better algorithms. Too many nested loops currently. Use benchmarking to compare different ideas.
- Caching. Normalise the input letters and have a small round-robin memory cache.
- Better error handling.
- Move the dictionary load before any fork()

## LICENSE

Wordfinder is free software; you can redistribute it and/or modify it under the same terms as perl itself.
