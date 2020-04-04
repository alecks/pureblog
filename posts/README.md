<img src="https://raw.githubusercontent.com/fjah/pureblog/master/assets/pureblog.png" width="10%">

![Go](https://github.com/fjah/pureblog/workflows/Go/badge.svg)

Pureblog is the most simple markdown blog you could possibly get. Written in Go using gin-gonic and golang-commonmark.

## Installing

The install process requires for you to have *Go*; get it from [here](https://golang.org/dl).

Clone or download the repository from [here](https://github.com/fjah/pureblog/archive/master.zip), then extract it. Open a terminal window (macOS/GNU+Linux) or a command prompt (Windows) in pureblog's directory. Follow by running the following command:
```bash
go get . && go build .
```

This will compile the program into a single executable. You can run this; use the `ls` command to view the name of it. It should be something along the lines of `pureblog` (macOS/GNU+Linux) or `pureblog.exe` (Windows). Simply use `./pureblog`, assuming that it's called that.

If you'd rather run it instantly without creating an executable, use the following:

```bash
go get . && go run .
```

## Configuring

Now that you have pureblog installed, you can configure it. The following directories and files can be safely changed.

### styles

Inside the folder you downloaded, you'll notice that there's a `styles` directory. This is where all of the CSS and HTML templates are kept. The comments at the start of the HTML files should help.

### .env

`.env` - the file in the repository's root directory - includes configurable values.

#### PORT

This variable is self-explanatory; it's the port on which portblog's HTTP server will listen. For example, if it's set to `8080`, the blog would be accessable by opening `http://localhost:8080`.

> Default: `80`

#### CACHE

This variable decides on whether or not posts will be cached. This provides quite a large performance improvement if you're working with large files.

**Note**: By default, pureblog caches *everything*, including the CSS, etc. If you want your site to update instantly without requiring a restart, change the `CACHE` value in `.env` to `FALSE`.

> Default: `TRUE`

## Running

If you've already compiled pureblog (as shown in the "Installation" chapter), running it is as simple as executing the generated file. You can test if the blog's working by visiting [here](http://localhost/README), assuming that the `PORT` variable in `.env` is set to `80`.