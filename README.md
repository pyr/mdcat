mdcat: terminal markdown cat
============================

A terminal utility to display markdown files.

## Building

```bash
go build
```

You can then use the resulting `mdcat` binary.

## Usage

```
mdcat [FILE]...
```

`mdcat` assumes all command-line arguments to be filenames to display on the
terminal. With no arguments, `mdcat` waits for input on **stdin** and produces
the parsed output once it is closed.

## Notes

`mdcat` is based on [go-termd](https://github.com/tj/go-termd), which does all
the work.

## Changelog

### 0.0.1

- Initial release

## Sandbox

Below you'll find a few typical markdown constructs, used to test
`mdcat`.

- One
- Two
- Three

```clojure
(defn some-function
  "docstring"
  [system foo]
  (do-bar-with-system system foo))
```

