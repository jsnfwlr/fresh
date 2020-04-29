# Refresh

Refresh is a command line tool that builds and (re)starts your web application everytime you save a Go or template file. It started life as [Fresh](https://github.com/gravityblast/fresh), which is now unmaintained.

If the web framework you are using supports the Refresh runner, it will show build errors on your browser.

It currently works with [Traffic](https://github.com/gravityblast/traffic), [Martini](https://github.com/codegangsta/martini), [Negroni](https://github.com/urfave/negroni) , and [gocraft/web](https://github.com/gocraft/web).

## Installation

```
go get github.com/jsnfwlr/refresh
```

## Usage

```
cd /path/to/myapp
refresh
```

Refresh will watch for file events, and every time you create/modify/delete a file it will build and restart the application.
If `go build` returns an error, it will log it in the tmp folder.

[Traffic](https://github.com/gravityblast/traffic) already has a middleware that shows the content of that file if it is present. This middleware is automatically added if you run a Traffic web app in dev mode with Refresh.
Check the `_examples` folder if you want to use it with Martini or Gocraft Web.

`refresh` uses `./refresh.conf` for configuration by default, but you may specify an alternative config filepath using `-c`:

    refresh -c other_runner.conf

Here is a sample config file with the default settings:

    root:              .
    tmp_path:          ./tmp
    build_name:        runner-build
    build_log:         runner-build-errors.log
    valid_ext:         .go, .tpl, .tmpl, .html
    no_rebuild_ext:    .tpl, .tmpl, .html
    ignored:           assets, tmp
    build_delay:       600
    colors:            1
    log_color_main:    cyan
    log_color_build:   yellow
    log_color_runner:  green
    log_color_watcher: magenta
    log_color_app:


## Authors

* [Jason Fowler](https://jsnfwlr.io) - as [Refresh](https://github.com/jsnfwlr/refresh)
* [Andrea Franz](http://gravityblast.com) - as [Fresh](https://github.com/gravityblast/fresh)

## More

* [Mailing List](https://groups.google.com/d/forum/golang-fresh)

## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request

