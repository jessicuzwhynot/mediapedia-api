[![Build Status](https://travis-ci.com/HammerMeetNail/mediapedia-api.svg?branch=master)](https://travis-ci.com/HammerMeetNail/mediapedia-api)
[![Coverage Status](https://coveralls.io/repos/github/HammerMeetNail/mediapedia-api/badge.svg?branch=master)](https://coveralls.io/github/HammerMeetNail/mediapedia-api?branch=master)
# mediapedia-api
An API for accessing media metadata

# Getting Started

```
$ docker build -t mediapedia-api:local .
$ docker-compose up -d
$ curl -X POST localhost/media/games
{"id":"5e3f4ddb2c9f1e0001c6403f","name":"Test","coverArt":"https://gamecoverart.com","description":"Lorem Ipsum","mediapediaRating":96,"metacritic":95,"publisher":"Microsoft","releaseDate":"2019-12-31","stores":[{"id":"","name":"Steam","url":"https://store.steampowered.com"},{"id":"","name":"GOG","url":"https://gog.com"}]}
$ curl localhost/media/games/5e3f4ddb2c9f1e0001c6403f
{"id":"5e3f4ddb2c9f1e0001c6403f","name":"Test","coverArt":"https://gamecoverart.com","description":"Lorem Ipsum","mediapediaRating":96,"metacritic":95,"publisher":"Microsoft","releaseDate":"2019-12-31","stores":[{"id":"","name":"Steam","url":"https://store.steampowered.com"},{"id":"","name":"GOG","url":"https://gog.com"}]}
```
