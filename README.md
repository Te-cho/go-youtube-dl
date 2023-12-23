# Go-Youtube-dl
A GoLang library and cli for downloading audio, video and subtitles with meta data using the known [Youtube-dl](https://rg3.github.io/youtube-dl/) application.

![Go-youtube-dl logo](https://golang.org/doc/gopher/appenginegopher.jpg "Golang Gopher")

---------------------------------------
* [Features](#features)
* [Requirements](#requirements)
* [Installation](#installation)
* [Usage](#usage)
* [License](#license)

---------------------------------------

## Features
* Usage of the famous known application [Youtube-dl](https://rg3.github.io/youtube-dl/) .
* Optional placeholder interpolation

## Requirements
* Go 1.2 or higher
* [Youtube-dl](https://rg3.github.io/youtube-dl/) application to be installed on your server or whatever the machine running the Go code.

---------------------------------------

## Installation
Simple install the package to your [$GOPATH](https://github.com/golang/go/wiki/GOPATH "GOPATH") with the [go tool](https://golang.org/cmd/go/ "go command") from shell:
```bash
$ go get github.com/youtube-videos/go-youtube-dl
```
Make sure [Git is installed](https://git-scm.com/downloads) on your machine and in your system's `PATH`.

## API Usage
Example of usage is as follow:
```bash
ytdl := youtube_dl.YoutubeDl{}
ytdl.Path = "$GOPATH/src/app/srts" // for example
err := ytdl.DownloadVideo(video.Id)
if err != nil {
    log.Printf("%v", err)
}
```
Then you can handle the downloaded file in the specified path.

## CLI Usage
```
go-youtube-dl -h
Usage of ./go-youtube-dl:
  -a    download audio.
  -f string
        Specific media format to download.
  -p string
        Path where you want to download. (default "/home/${USER}/Music")
  -s    Download subs from the youtube for the specified video id
  -v    download video.
  -y    Find where youtube-dl is installed.
  -yid string
        ID of the youtube video to download subs
```
---------------------------------------

## License
Go-Youtube-dl is licensed under the [Mozilla Public License Version 2.0](https://raw.github.com/go-sql-driver/mysql/master/LICENSE)

Mozilla summarizes the license scope as follows:
> MPL: The copyleft applies to any files containing MPLed code.


That means:
* You can **use** the **unchanged** source code both in private and commercially
* When distributing, you **must publish** the source code of any **changed files** licensed under the MPL 2.0 under a) the MPL 2.0 itself or b) a compatible license (e.g. GPL 3.0 or Apache License 2.0)
* You **needn't publish** the source code of your library as long as the files licensed under the MPL 2.0 are **unchanged**

Please read the [MPL 2.0 FAQ](https://www.mozilla.org/en-US/MPL/2.0/FAQ/) if you have further questions regarding the license.

