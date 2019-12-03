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

## Send Queue
Send Queue
```
./main send -f templates/send/duck.yml
```

or create your project directory then see templates/send/duck.yml for value

```
./main send
```

## Get Result Queue
Get Result Queue

```
./main get -i $UUID -f templates/send/duck.yml
```