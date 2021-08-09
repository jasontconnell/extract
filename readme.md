*** Extract

Extract is a text content grep across a whole directory

Usage:

```extract -d . -e txt "^First Name: (.*)$" "^Last Name: (.*)$" "Found name: %s %s"```

It will output one line per match