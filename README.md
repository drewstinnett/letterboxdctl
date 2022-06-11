# letterboxdctl

## Overview

Interact with letterboxd.com from the command line.

## Installation

Download from the [GitHub releases page](https://github.com/drewstinnett/letterboxdctl/releases)

Use Homebrew to install:

```shell
$ brew tap drewstinnett/tap
...
$ brew install letterboxdctl
...
```

## Usage

Check out the help page for full info:

```shell
$ letterboxdctl --help
...
```

A few examples:

```shell
❯ letterboxdctl watchlist mondodrew
...
- id: "635254"
  title: Lovers Rock
  slug: lovers-rock-2020
  target: /film/lovers-rock-2020/
  year: 2020
  externalids:
    imdb: tt10551102
    tmdb: "712166"
- id: "503640"
  title: New Order
  slug: new-order-2020
  target: /film/new-order-2020/
  year: 2020
  externalids:
    imdb: tt12474056
    tmdb: "575446"
   • Stats                     duration=3.980591798s total=78
```
