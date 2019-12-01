# DUCK
Queue Task Schedule service

## Build
Build duck
```
go build src/main.go
```

or execute makefile

```
make
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