# caicloud-formatting

## Getting Started

Suppose you have `caicloud-formatting` and `config-admin` in your GOPATH and both of them are up to date, and you are currently in the root of the latter repository.

1. Simply run `go run $GOPATH/src/github.com/hezhizhen/caicloud-formatting/main.go` and boom! Everything's done.
2. Add a Makefile target which actually executes the command above, and `make` it.

```makefile
FORMATTING := $(GOPATH)/src/github.com/hezhizhen/caicloud-formatting
format:
    @go run $(FORMATTING)/main.go
```

## imports

Adjust the imports order to satisfy our [convention](https://github.com/caicloud/engineering/blob/master/guidelines/golang.md#order).

### Usage

You can directly run the `main.go`, or you can specify the root, like `go run $GOPATH/src/github.com/hezhizhen/caicloud-formatting/main.go .`. Running it in the root of a repository is highly recommended, because files in the vendor should be skipped, but they are identified by checking if paths have prefixes of `vendor`.

A file will be skipped if it is:

* in the vendor directory
* a hidden one
* not a Go file (end with `.go`)

Besides, commented import lines will be removed to make imports clean.

## Git integration

Here's an idea about how to use `git pre-commit` to automatically format our code.

1. Install binary from `hezhizhen/caicloud-formatting`
2. Add a `pre-commit` file to `.git/hooks` directory and make it executable
3. Add and commit as usual

### Install binary

1. Get the latest code: `go get -u -v github.com/hezhizhen/caicloud-formatting`
2. Enter the repo: `cd $GOPATH/src/github.com/hezhizhen/caicloud-formatting`
3. install it: `go install ./...`

### Add `pre-commit`

Save the following code to a file named `pre-commit` in the `.git/hooks` directory:

```sh
#!/bin/sh

caicloud-formatting .
if [[ -n $(git diff) ]]; then
    echo "Imports have been changed. Please commit again."
    exit 1
fi
exit 0
```

Don't forget to make it executable: `chmod +x .git/hooks/pre-commit`, Or you'll see such messages:

```text
hint: The 'pre-commit' hook was ignored because it's not set as executable.
hint: You can disable this warning with `git config advice.ignoredHook false`
```

NOTE: there is a file `install` in the root which will do the adding stuff automatically by simply running `bash install` in the root of your repository.

1. `wget https://raw.githubusercontent.com/hezhizhen/caicloud-formatting/master/install`
2. `bash install`
3. `rm install`

### (Optional) Configure git hooks path

To share git hooks with your team, it's better to configure a different path of git hooks `.githooks` instead of using the default one `.git/hooks`.

```sh
git config core.hooksPath .githooks
```

Then add the `pre-commit` hook to the directory. Maybe you can also put the binary here (If so, don't forget to change the path of `caicloud-formatting` in `pre-commit`).

