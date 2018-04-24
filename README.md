# FemptoServer

FemptoServer is a tiny server to be able to test an HTML (or other) folder in your browser.

## Usage

Easy!

#### To compile

```bash
$ git clone 
$ cd femptoserver
$ go build serve.go
$ sudo cp serve /bin
```

#### To execute

```bash
$ cd [your folder]
$ serve [--port=NNNN]
```

Default port is 5555.

License MIT, 2018