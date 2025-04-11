# serve.lee.io

Simple HTTP server for serving files/directories

## Usage

* Default (port `8080`, current directory):

```
curl serve.lee.io | sh
```

* Specify port:

```
curl serve.lee.io | sh -s -- --port 8123
```

* Specify directory:

```
curl serve.lee.io | sh -s -- --directory /tmp/test
```
