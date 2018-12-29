# GitHub Stars

Lists your most recent GitHub stars in a compiled list that can be sent in a
digest format.

    $ gh-stars
    jez/pandoc-starter
    CDSoft/pp
    haya14busa/reviewdog

Stars are listed in a "week" basis starting "Sunday".

	$ gh-stars --help
	Usage: gh-stars [options]
	Options:
	-h, --help           Displays this help message
	-i, --insecure       Disables HTTPS certificate verification
	-l, --limit int      Limit the amount of starts retrieved
	-t, --token string   A GitHub Personal API Access Token
	-v, --version        Displays version information
	https://github.com/augustohp/gh-stars

You can also set `GHSTARS_GITHUB_TOKEN` environment variable instead allowing
your GitHub token lurking in your history.

## Development

There are [make][] targets for:

* `build` compiles the project and puts it in your `$GOTAH/bin`
* `release` creates binaries for *Windows*, *Linux* and *MacOS* under `dist/`

There are no *automated tests* and the code lacks love.

[make]: https://www.gnu.org/software/make/
