*** Extract

Extract is a text content grep across a whole directory

Usage:

```extract -d . -e txt -r "First Name: (.*)" -o "Found first name: %s"```

It will output one line per match