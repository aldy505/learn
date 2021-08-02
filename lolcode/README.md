# LOLCODE

This is an esoteric programming language that I've heard since the early 2010s and I understood it from the beginning.

The documentation for the language: https://github.com/justinmeza/lolcode-spec/blob/master/v1.2/lolcode-spec-v1.2.md

Another good resource to learn it quickly: https://github.com/leyarotheconquerer/lolcode-reference/blob/master/Reference.md

## Installation

Simply do these:
```bash
$ git clone https://github.com/justinmeza/lci.git
$ cd lci
$ cmake .
$ make
$ sudo make install
# delete the lci folder if you want
$ cd .. && rm -rf lci
```

Check if the installation is complete:
```sh
$ lci -v
lci v0.10.5
```

## How to run the files

Say you're in the `basic` directory. Simply do this:
```sh
$ lci conditional.lol
# results goes here
```