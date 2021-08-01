# 🍄 Mycorrhiza Wiki
<img src="https://mycorrhiza.wiki/binary/release/1.3/screenshot" alt="A screenshot of Mycorrhiza Wiki's home page in the Safari browser" width="600">

**Mycorrhiza Wiki** is a filesystem-backed wiki engine that uses Git for
keeping history.

[👉 Main wiki](https://mycorrhiza.wiki)

## Features

* **No database required.** Everything is stored as simple files. It makes installation super easy, and you can modify the content directly.
* **Everything is hyphae.** A [hypha][feature-hypha] is a unit of content such as a picture, video or a text article. Hyphae can [transclude][feature-transclusion] and link each other resulting in a tight network of hypertext pages.
* **Hyphae are authored in [Mycomarkup][mycomarkup],** a custom markup language that's designed to be unambigious and easy to use.
* **Nesting of hyphae** is supported. A tree of related hyphae is shown on every page.
* **History of changes** for textual parts of hyphae. Every change is safely stored in [Git][integration-git]. Web feeds for recent changes included!
* **Keyboard-driven navigation.** Press `?` to see the list of shortcuts.
* **Support for [authorization][feature-authorization].**
* **[Opengraph][standard-og] support.**
* **Optional [Telegram][telegram] integraion.**

[feature-hypha]: https://mycorrhiza.wiki/hypha/feature/hypha
[feature-transclusion]: https://mycorrhiza.wiki/hypha/feature/transclusion
[feature-authorization]: https://mycorrhiza.wiki/hypha/feature/authorization
[mycomarkup]: https://mycorrhiza.wiki/help/en/mycomarkup
[integration-git]: https://mycorrhiza.wiki/hypha/integration/git
[standard-og]: https://mycorrhiza.wiki/hypha/standard/opengraph
[telegram]: https://mycorrhiza.wiki/help/en/telegram

## Installing

See [the deployment guide][deployment] on the wiki.

[deployment]: https://mycorrhiza.wiki/hypha/guide/deployment

## Contributing

Help is always welcome! We have an IRC channel [#mycorrhiza on
irc.libera.chat][irc] and a [Telegram chat][tg] for discussions and
development. You can also sponsor the maintainer of Mycorrhiza,
[@bouncepaw][bp], on [Boosty][bp-donate]. If you want to contribute with code,
you can either open a pull request on GitHub or send a patch to the [mailing
list][mlist]. Feel free to open an issue on GitHub or contact us directly.

[irc]: irc://irc.libera.chat/#mycorrhiza
[tg]: https://t.me/mycorrhizadev
[bp]: https://github.com/bouncepaw
[bp-donate]: https://boosty.to/bouncepaw
[mlist]: https://lists.sr.ht/~handlerug/mycorrhiza-devel

You can view the list of planned features on [our GitHub project kanban
board](https://github.com/bouncepaw/mycorrhiza/projects/1).

Compare Mycorrhiza Wiki with other engines on [WikiMatrix](https://www.wikimatrix.org/show/mycorrhiza).