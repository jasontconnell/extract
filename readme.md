*** Extract

Extract is a text content grep across a whole directory

Usage:

```extract -d . -e txt "^First Name: (.*)$" "Found name: %s"```

Or you can search and output multiple matches. It will parse a `regex output` format, for each pair of supplied inputs. Each regex must have a corresponding output.

```extract -d . -e log "Problem with ID ([0-9a-f\-]+)" "Problem ID: %s" "Error message: \[(.*?)\]" "Error message: %s"```



It will output one line per match