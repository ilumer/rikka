# Rikkac - CLI tool of Rikka

[中文版][version-zh]

Rikkac need to be used with a [Rikka][rikka] server.

## Usage

`rikkac <format> filename`

`<format>` can be:

- `-s`: Src, image source url
- `-m`: Markdown
- `-h`: HTML
- `-b`: BBCode
- `-r` reStructuredText

Src is default format. Format priority as same as the list above, lowest to highest. This is, `-m -b` considered as `-b`, `-m` is ignored. Not so complicated, you shouldn't remember priority if you never provide two format in one command.

## Build and Install

### Executable Binary Download

Now we only provide [executable binary for Linux][download], Because I only have Linux installed in my PC, QwQ

Then rename the file to `rikkac` and move to a folder in your `PATH`.

OK, installation finished, now you need [configure](#configure_and_usage) Rikkac before use it.

User of other os please refer to next section to build and install Rikkac.

### From Source Code

First, you need have Golang installed in your PC, then:

`go get github.com/7sDream/rikka/rikkac`

Add `$GOPATH/bin` into your `PATH`, if you haven't do this when you install Golang.

Then run `rikkac --version`, a version number means install successfully.

You need some [configure](#configure-and-usage) before use Rikkac.

## Configure and Usage

Rikkac need to env variable： `RIKKA_HOST` and `RIKKA_PWD`. for  Rikka server address and password.

![](http://7sdream-rikka-demo.daoapp.io/files/2016-09-05-066558195)

Than you can enjoy Rikkac.

Just run `rikkac -m filepath` for upload.

You can get detail log when you meet some error by add  `-v` or `-vv` option.

## Multi File upload

Just provide file path one by one: `rikkac -m file1 file2 file3 ...`

Or you can use wildcard if your shell support：`rikkac -m *.png`。

![](http://odbw8jckg.bkt.clouddn.com/ba2d2dca-2ae2-4436-ade2-7905183ce23d.png)

## Tips: Copy Result to Clipboard in Quick

![](http://7sdream-rikka-demo.daoapp.io/files/2016-09-05-781037494)

need xclip installed：`apt-get install xclip`.

[version-zh]: https://github.com/7sDream/rikka/blob/master/rikkac/README.zh.md

[rikka]: https://github.com/7sDream/rikka
[download]: https://github.com/7sDream/rikka/releases/tag/Rikkac
