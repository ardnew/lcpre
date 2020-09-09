# lexroot
##### Determine lexical root of multiple strings


`lexroot` finds the lexical root, or, the longest common prefix, of a given set of strings. The components of each string may be partitioned into substrings separated by a common delimeter substring.

For example, if the given strings are file paths and the delimeter is the file path separator (e.g., `/`), it will find the nearest common parent directory:

```
$ lexroot -s "/" "/R/P/A/A1" "/R/P/B" "/R/P/C/C1/C1.1"
/R/P

$ lexroot -s "/" "/A/B/C" "/D/E/F"
/

$ lexroot -s "/" "a/b/c" "d/e/f"


$ lexroot -s "/" "a/b/c" "a/d"
a
```

By default, the delimeter is the file path separator of the OS. So on a Unix system, the `-s "/"` argument could have been left off in all of the examples above.

If a single string is input, it is returned verbatim. If no strings are input, or there does not exist a common prefix, the empty string is returned.
