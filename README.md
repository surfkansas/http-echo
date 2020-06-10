# http-echo
The simple tool echos HTTP headers and environment variables back in an text result.

It is useful when trying to debug the environment and headers that are interacting with a website. 

Note: since this echos all environment variables - including any secrets - it should NEVER be used in a production website.

# Usage

Running the following command will start up a simple version of the container to demonstrate how it works:

```
docker run -p:8080:8080 surfkansas/http-echo
```