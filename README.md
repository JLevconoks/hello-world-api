A simple Hello, World! API for some experiments with Kubernetes.

Endpoints "/health", "/ready/ and "/hello".

Startup delay and port can be configured via enviroment variables:
- `STARTUP_DELAY` in seconds, defaults to 0
- `PORT` defaults to 8080 

```bash
docker run --name hello-api \
-p 8080:8080 \
-e STARTUP_DELAY=10 \
jlevconoks/hello-world
```