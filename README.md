# 🍄 MycorrhizaWiki 0.9
A wiki engine.

## 0.9
This is a development branch for 0.9 version. Features I want to implement in this release:
* [x] Recent changes page.
* [x] Hypha deletion.
* [ ] Hypha renaming.
* [x] Support async git ops.

## Installation
```sh
git clone --recurse-submodules https://github.com/bouncepaw/mycorrhiza
cd mycorrhiza
make
# That make will:
# * run the default wiki. You can edit it right away.
# * create an executable called `mycorrhiza`. Run it with path to your wiki.
```

## Current features
* Edit pages through html forms
* Responsive design
* Works in text browsers
* Wiki pages (called hyphae) are in gemtext
* Everything is stored as simple files, no database required
* Page trees
* Changes are saved to git
* List of hyphae page
* History page
* Random page
* Light on resources: I run a home wiki on this engine 24/7 at an [Orange π Lite](http://www.orangepi.org/orangepilite/).

## Contributing
Help is always needed. We have a [tg chat](https://t.me/mycorrhizadev) where some development is coordinated. Feel free to open an issue or contact me.

## Future plans
* Tagging system
* Authorization
* Better history viewing
* More markups
