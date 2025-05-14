# Bye Bye, Link Tracking 👋

> [!IMPORTANT]
> This software is still in early development. See planned features.

Are you a part-time privacy preserver and want to bring your link-sharing to the next level? _Bye bye, link tracking_ has your back (well, actually your link). Because who needs link tracking anyway.

## What does it do?

You probably stumbled across `si` in YouTube links (e.g. <https://youtu.be/xvFZjo5PgG0?si=SomeFancyTextHere>) or similar add-ons to links (e.g. [UTM parameters](https://en.wikipedia.org/wiki/UTM_parameters)) at some time already. These additions are not necessary for any functionality but rather allow for tracking link clicks, marketing campaigns, etc.

_Bye bye, link tracking_ is installed as a service on your system that monitors your clipboard. As soon as it identifies a URL in your clipboard that has tracking additions, it automatically replaces the link with a purged version. You manually removing it is not necessary anymore.

![Gandalf "You shall not pass" meme adopted to "You shall not track my links"](/img/gandalf.jpeg)

## How to Install

Download a release from the releases tab and move the config file to `/home/$USER/.config/byebyelinktracking/config.json`.

If you want to autostart `byebyelinktracking` for your user, move it to `/usr/bin/byebyelinktracking` and create a file `byebyelinktracking.service` at `/home/$USER/.config/systemd/user/` with the following content (file is also available in the repository):

```
[Unit]
Description=Removes tracking from copied links.

[Service]
ExecStart=/usr/bin/byebyelinktracking
Restart=always

[Install]
WantedBy=multi-user.target
```

> [!NOTE]
> If you have a custom config location (not `/home/$USER/.config/byebyelinktracking/config.json`), you can add the `-c` option at `ExecStart`.

Then run the following commands to enable your service:

```sh
systemctl --user daemon-reload
systemctl --user enable byebyelinktracking.service
```

You're good to go!

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

- [x] YAML Support for config files
- [x] Service installation instructions
- [ ] Further code tidying and more options depending on necessity
- [ ] Muuuuuch more entries to the config file :) (PRs welcome)
