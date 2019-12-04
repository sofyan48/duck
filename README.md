# DUCK
Queue Task Schedule service

## Build
Build duck
```
go build duck src/main.go
```

or execute makefile

```
make build
```

## Setting environment for duck
Create .duck file in home user directory and see env.example for value

```
nano ~/.duck
```


## Start Worker
Start Worker
```
./main worker start --name=duckTest --f templates/worker/task.yml
```
## Queue CLI Mode
### Send Queue
Send Queue
```
./main send -f templates/send/duck.yml
```

or create your project directory then see templates/send/duck.yml for value

```
./main send
```

### Get Result Queue CLI MODE
Get Result Queue

```
./main get -i $UUID -f templates/send/duck.yml
```

## Queue REST Mode
### Starting Server
```
./main server start
```

### Send Queue
```
curl -X POST \
  http://localhost:5000/api/send/task \
  -H 'Content-Type: application/json' \
  -H 'cache-control: no-cache' \
  -d '{
  "duck": {
    "task": "Send To Server GET",
    "action": {
      "url": "$URL", ## URL DEFINE
      "method": "GET",
      "trigger": "request",
      "worker": "$WORKER" ## WORKER NAME
    },
    "headers": [
      {
        "name": "Authorization",
        "value": "xxxx",
        "type": "string"
      }
    ],
    "parameter": [
      {
        "name": "microsite_url",
        "value": "event-name-2740",
        "type": "string"
      }
    ],
    "setting": {
      "loop": true
    }
  }
}'
```

### Get Result
```
curl -X POST \
  http://localhost:5000/api/get/task \
  -H 'Content-Type: application/json' \
  -H 'cache-control: no-cache' \
  -d '{
	  "duck": {
	  	"uuid" : "c9ab13c9410251c4c0f83d9b",
	    "action": {
	      "trigger": "request",
	      "worker": "TEST1"
	    }
	}
}'
```