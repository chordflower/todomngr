# PKGMAN

This is a simple package manager for windows, using the user local home.


## Features

- Support for remote repositories;
- User based installs;
- Made to be a companion for cmder;
- Easy to use.


## Installation

Install pkgman with go get

```bash
  go install -u github.com/chordflower/pkgman/cmd/pkgman
```

## Usage/Examples

### Install a package

```bash
pkgman install pkgname
```

### List installed packages

```bash
pkgman List
```

### Search installed package by expression

```bash
pkgman search --local expression
```

### Search remote package by expression

```bash
pkgman search --remote expression
```

### Remove a package

```bash
pkgman remove pkgname
```

### Get information about a package

```bash
pkgman info pkgname
```

### Update the installed packages

```bash
pkgman update
```

### Update a specific package

```bash
pkgman update pkgname
```

### Add repository

```bash
pkgman repositories add --name="remote name" "remote_path"
```

### List repositories

```bash
pkgman repositories list
```

### Remove repository

```bash
pkgman repositories remove --name="remote name"
```

### Update repositories

```bash
pkgman repositories update
```
## Authors

- [@carddamom](https://www.github.com/carddamom)


## License

[Apache-2.0](https://choosealicense.com/licenses/apache-2.0/)


## Contributing

Contributions are always welcome!

See `contributing.md` for ways to get started.

Please adhere to this project's `code of conduct`.


## Roadmap

- Add commands to support the creation of remote repositories;
- Add support for ssh repositories;
- Add support for webdav repositories.

## Acknowledgements

 - [Awesome GO from Avelino](https://github.com/avelino/awesome-go) => For its awesome list of golang libraries;
 - [pkg.go.dev](https://pkg.go.dev) => For reading documentation about go packages;
 - [Readme.so](https://readme.so/editor) => For creating this readme file;
 - [Scoop](https://scoop.sh/) => For inspiration for this program.
