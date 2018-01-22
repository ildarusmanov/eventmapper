Event mapper service
==========================

[![Build Status](https://travis-ci.org/ildarusmanov/eventmapper.svg?branch=master)](https://travis-ci.org/ildarusmanov/eventmapper)

## Setup

1. Install and run RabbitMQ
2. Install Docker
3. Clone the repo
4. Run the following commands:

```
cd eventmapper
sudo docker build -t eventmapper .
// prod
sudo docker run --restart=always -d -p 10.90.137.73:8000:8000 --network host eventmapper
// or dev
sudo docker run -p 8000:8000 --network host eventmapper 
// list containers
sudo docker ps
```

## Send data

```
# POST /{queue_key}/events?token={token}
# e.g. POST /mysite/events?token=super-token
{
  "source": {
    "source_id": "example.com",
    "source_type": "http",
    "origin": "/api/v1/endpoint/1.json",
    "params": {
      "format": "json"
    }
  },
  "target": {
    "target_type": "User",
    "target_id": "1",
    "params": {
      "email": "test@email.com"
    }
  },
  "event_name": "authorized",
  "user_id": "1",
  "created_at": 1712311,
  "params": {
    "userAgent":"Mozilla\/5.0 (X11; Ubuntu; Linux x86_64; rv:54.0) Gecko\/20100101 Firefox\/54.0",
    "userIP":"127.0.0.1"
  }
}

```

## Run tests
```
# update config_test.yml
cd {project_directory}
go test ./models ./controllers ./mq
```

## Todo

* Increase test coverage