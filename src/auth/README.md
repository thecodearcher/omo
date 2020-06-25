# 🔐Auth Service 

🚧Everything relating to auth will eventually be here.

## How to 🏃🏼‍♂️
- Point your terminal to this directory `src/auth`
- Build Docker image
``` 
docker build -t omo-auth-service .  
```
- Run Docker image
```
docker run -it --rm -p 8080:8080 -v $PWD:/go/src/auth omo-auth-service
```
> _**Note:** This will add your current directory as a volume to the container so ensure you are in `src/auth`_