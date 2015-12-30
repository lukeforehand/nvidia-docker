
## build DIGITS container
```
docker build -t digits samples/ubuntu/digits
```

## run DIGITS container
```
./nvidia-docker run -p 34448:34448 digits &
docker ps
```

## to login to running DIGITS container
```
docker exec -it $(docker ps | grep digits | cut -d' ' -f1) bash
```

## access DIGITS interface
```
curl localhost:34448
```

## kill DIGITS container
```
docker kill $(docker ps | grep digits | cut -d' ' -f1)
```

