# opensecrets golang API

This is a golang API client for [opensecrets.org](http://opensecrets.org)

## Example usage

See [cli/main.go](https://github.com/spudtrooper/opensecrets/blob/main/cli/main.go)

### Deployed frontend

https://api-dumpster-fire.herokuapp.com/opensecrets

## CLI

The binary associated with this module is a minimal CLi which serves as an example use case.

To install:

```
$ go install github.com/spudtrooper/opensecrets
```

See the available commands:

```
$ opensecrets
```

Example call for John Kennedy of LA:

```
$ opensecrets gl --id N00026823
```


Example call for all candidates in LA

```
$ opensecrets gl --id LA
```