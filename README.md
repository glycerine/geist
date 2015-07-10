# geist : golang scripting

# Note: 2015/July/09: Currently complete vaporware: this is a vision statement
rather than a working repository.

> Geist (German pronunciation: [ˈɡaɪst]) is a German word.
> Depending on context it can be translated as the English
> words mind, spirit, or ghost, covering the semantic field
> of these three English nouns.
>
>    -- https://en.wikipedia.org/wiki/Geist

The problem with most scripting languages is that
they are not actually great programming languages.

But since they are convenient, quick to write,
and have no compilation overhead, they get used.

Alot.

That's not a good thing.

Because: often times a script grows up and becomes
a big ball of mud: a truly hard to maintain program.

Good testing and software engineering becomes really difficult,
and one faces a full rewrite of the overgrown
script in another language.

In really bad cases, I've seen shops expend years of effort by multiple people on such rewrites.

Ideally a script that has grown up should be re-written in a
language with good tooling, static type checking,
and good performance.

Geist puts forward this thesis: Why not write
in that eventual language *from the start*?

The geist script-runner turns Go (golang) into
a scripting language: the one you should write
your scripts in from the start.

# geist: do it right from the start. script in Go.

Geist makes it viable to write your scripts in
Go, and then when they need to grow into actual
programs, the transition is a seemless matter
of adding a few lines of boilerplate to your
script and moving it to a Go source heirarchy.

More: your configuration files can be written
as geist. Go has fantastic support for declaring
literal values, and this makes configuration easy. No more need
to invent a new configuration file format for
every new program.

Instead of inventing another scripting language,
or another configuration format, geist offers a way
to use Go as a scripting and configuration language
from day one.

A geist script is #! runnable, there is no separate
manual compilation step. Not even "go run" is needed
on your part. Geist takes care of building and caching
your script as an executable. You will need a
go installation to get and build geist and geist-install,
however once geist is installed with 'geist install',
geist takes care of isolating itself from any other GOROOT or GOPATH.

The geist executable will read the geist script, and
will add a couple of lines of boilerplate that turn it into
a real Go program. e.g. geist adds the "package main"
declaration, some imports, and wraps your lines in
the main() function.

Then geist will compile and run it for you.
The compilation actually only happens if the script
has been changed (we cache the prior version); so
geist offers even faster startup time than a
"go run" invocation after the first compile.

Compilation happens in an isolated GOROOT/GOPATH
environment under $HOME/.geist by default (settable
with GEIST_HOME env variable). This means you can
run geist starting in any directory, and it will
ignore any other golang code, GOROOT/GOPATH settings.

A given geist script has an implicit main
and package use is automatically detected to the
extent possible (by goimports). "geist-install"
acts as the "go install" replacement to install
a new package under your ~/.geist/src heirarchy.
The "geist-install" tool is used to initialize
your $HOME/.geist directory as well. It need only
be run once, without arguments, to do this.

Additional imports can be added anywhere, no problem.
geist will collect them all to the top of the file
before compiling it.

Natually structs and functions can be defined in a geist
file, and the implicit main stops at the point
of the first (if any) struct or fuction definition.
the //endmain comment can also be used to terminate
the automatically added main() function.

Planned additions: 

 * ```//include <path>``` will, in the future, be used to include additional
geist (or .go) files from the main script. If the
file is a .go file, then the package declaration
will need to be removed and the imports merged.

 * For configuration file support, and DSL
replacement support, we'll want a means
of including a user-defined and pre-defined
set of functions/structs in the main package that
all such local geist scripts should have
access to.
