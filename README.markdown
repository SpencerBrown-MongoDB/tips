# Setting up go to access private repos with dependendent projects

Normally, Go resolves dependencies by going to the public golang.org site. If you have a dependency that's in a private GitHub repo, here's how to do it. 

Set the GOPRIVATE Go Environment Variable to access `github.com/<org>/<repo>` as a private library. You can also use a wildcard `*` for the repo, to designate any repo in that organization. 

```bash
go env -w GOPRIVATE=github.com/<org>/<repo-or-asterisk>
```

## Managing a Go package that others depend on

* Use semver to mark versions, e.g. v0.9.5. The leading "v" is required.
* Tag a package release with the semver on GitHub. 
* Create a GitHub Release for that version tag. 

## Testing with your private package


If you want to develop locally against a private package:

1. Clone the private package repo so its directory is "next to" the referring repo locally.
2. Add this to the `go.mod` of the referring repo:

```
replace github.com/<org>/<repo> => ../<repo>
```

When you push the private package changes into a new release, you can comment out that "replace" line in go.mod. 


# Setup for `gh` automatic authentication

1. Get a GitHub personal access token. 
2. Install the `gh` command (GitHub CLI). For Mac, `brew install gh`
3. Run `echo '<the-token>' | gh auth login --with-token

If you need to access another org via SSO:

* Run `gh release view -R <org>/<repo>` for one of the org's private repos. 
* It will prompt you to authorize your personal access token to the other organization via SSO. You can do this through the GitHub website also if you wish.

# Private fork of GitHub public repo

As GitHub does not natively support forking a public repo into a private repo, we have to do it manually. The procedure is:

## Duplicate/mirror the public repository to your private repository

see [the GitHub docs](https://docs.github.com/en/repositories/creating-and-managing-repositories/duplicating-a-repository)

create the empty private repo on GitHub, then:

```bash
cd /tmp
git clone --bare https://github.com/<current-org>/<public-repo>
cd <public-repo>.git
git push --mirror https://github.com/<new-orgorg>/<private-repo>
cd ..
rm -rf <public-repo>.git
```

Now you have a copy of the public repo in your private repo. To sync changes from the public repo into the private repo, create a remote for the public repo:

```bash
git clone https://github.com/<new-org>/<private-repo>
git remote add upstream https://github.com/<old-org>/<public-repo>
git fetch upstream master
git pull upstream master
```

Note this is not an official GitHub forked repository and you cannot initiate PRs from it. You would have to have an actual public fork to submit PRs. 
