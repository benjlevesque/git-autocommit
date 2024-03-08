# git-autocommit

Generates a [conventional commit message](https://www.conventionalcommits.org/en/v1.0.0/) based on the current branch name.


## Install
```bash
curl -sSfL https://raw.githubusercontent.com/benjlevesque/git-autocommit/main/godownloader.sh | sh -s -- -b /usr/local/bin
git config --global alias.autocommit "!git commit -m \"$(git-autocommit)\""
```

You can change the installation path by editing `/usr/local/bin` in the above command. Make sure the directory is in your `$PATH`

## Usage
1. Create a branch name based on your work (see [Examples](#Examples)):
```bash
gcb feat/users/create-user
```
> Tip: `gcb` is an alias for `git checkout -b`

2. Stage your files
```bash
git add main.go
```

3. Commit! 
```bash
git autocommit
```

The generated commit will be `feat(users): create user`


## Examples

| Branch name            | Message                  |
| ---------------------- | ------------------------ |
| feat/users/create-user | feat(users): create user |
| chore/cleanup          | chore: cleanup           |
| fix/foo/hello--world   | fix(foo): hello-world    |


## See also
- [git-branch-rename](https://github.com/benjlevesque/git-branch-rename): rename your current branch using your editor
- [gh](https://github.com/cli/cli): Github CLI
