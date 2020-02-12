[![Build Status](https://travis-ci.com/HammerMeetNail/mediapedia-api.svg?branch=master)](https://travis-ci.com/HammerMeetNail/mediapedia-api)
[![Coverage Status](https://coveralls.io/repos/github/HammerMeetNail/mediapedia-api/badge.svg?branch=master)](https://coveralls.io/github/HammerMeetNail/mediapedia-api?branch=master)
# mediapedia-api
An API for accessing media metadata

# Getting Started

```
$ docker build -t mediapedia-api:local .
$ docker-compose up -d
$ curl http://localhost/media/games/1?pretty
{
  "id": "1",
  "title": "Bloodborne",
  "coverArt": "https://gamecoverart.com",
  "description": "Lorem Ipsum",
  "mediapediaRating": 96,
  "metacritic": 95,
  "publisher": "FromSoftware",
  "releaseDate": "2019-12-31T00:00:00Z"
}
```
