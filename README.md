# Bye Bye, Link Tracking 👋

> [!IMPORTANT]
> This software is still in early development. See planned features.

Are you a part-time privacy preserver and want to bring your link-sharing to the next level? _Bye bye, link tracking_ has your back (well, actually your link). Because who needs link tracking anyway.

## What does it do?

You probably stumbled across `si` in YouTube links (e.g. <https://youtu.be/xvFZjo5PgG0?si=SomeFancyTextHere>) or similar add-ons to links (e.g. [UTM parameters](https://en.wikipedia.org/wiki/UTM_parameters)) at some time already. These additions are not necessary for any functionality but rather allow for tracking link clicks, marketing campaigns, etc.

_Bye bye, link tracking_ is installed as a service on your system that monitors your clipboard. As soon as it identifies a URL in your clipboard that has tracking additions, it automatically replaces the link with a purged version. You manually removing it is not necessary anymore.

![Gandalf "You shall not pass" meme adopted to "You shall not track my links"](/img/gandalf.jpeg)

## How to Install

Download a release from the releases tab

### Build from Source

Clone this repository

```sh
git clone https://github.com/niri81/byebyelinktracking
```

Build the Executable

```sh
go build -o byebyelinktracking
```

Run the Executable

```sh
./byebyelinktracking -h
```

## Planned Features

- [ ] YAML/TOML Support for config files
- [ ] Service installation instructions
- [ ] Further code tidying and more options depending on necessity
- [ ] Muuuuuch more entries to the config file :) (PRs welcome)
