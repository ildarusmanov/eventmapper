Event mapperservice
==========================
## Setup

1. Install and run RabbitMQ
2. Install Docker
3. Clone the repo
4. Run the following commands:

```
cd eventmapper
sudo docker build -t eventmapper .
sudo docker run -d -p 8000:8000 --network host eventmapper
sudo docker ps
```

## Save some log data

```
# POST /create/{queue_key}?token={token}
# e.g. POST /create/mysite?token=super-token
{
   "EventName":"view-page",
   "EventTarget":"sdfsd",
   "UserId":"1",
   "CreatedAt":1503489779,
   "Params":{
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