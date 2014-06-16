# Covergen ![Coverage](http://img.shields.io/badge/coverage-50.0%25-red.svg)


Autoupdate a coverage badge in your README.md with this one dumb trick!

To try it out:

```shell
go get github.com/natefinch/covergen
```

Uncomment the commented out code in foo_test.go.  Then rerun go test:

```shell
go test
```

Now check out the README.md ... it is updated automagically!

The only drawback is that you more than double the time it takes to run your tests...

Check out covergen_test.go for how it works.

