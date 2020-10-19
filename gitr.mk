
# git general stuff

# hardcoded
GITR_SERVER ?= github.com
GITR_ORG_UPSTREAM ?= getcouragenow

# reflected
GITR_ORG_FORK=$(shell basename $(dir $(abspath $(dir $$PWD))))
GITR_USER=$(GITR_ORG_FORK)
GITR_REPO_NAME=$(notdir $(shell pwd))
GITR_LAST_TAG=$(shell git describe --exact-match --tags $(shell git rev-parse HEAD))

# calculated
# upstream
GITR_REPO_UPSTREAM_ABS_URL=https:///$(GITR_SERVER)/$(GITR_ORG_UPSTREAM)/$(GITR_REPO_NAME)

GITR_REPO_ABS_URL=https:///$(GITR_SERVER)/$(GITR_ORG_FORK)/$(GITR_REPO_NAME)
GITR_REPO_ABS_FSPATH=$(GOPATH)/src/$(GITR_SERVER)/$(GITR_ORG_FORK)/$(GITR_REPO_NAME)

# remove the "v" prefix
GITR_VERSION ?= $(shell echo $(TAGGED_VERSION) | cut -c 2-)
#GITR_BRANCH_NAME=main # Later for gitea
GITR_BRANCH_NAME=master


### Variables passed in by dev

# EG: make M='a' T='v1.1.1' gitr-print

# M for commit message
#GITR_COMMIT_MESSAGE ?= autocommit
GITR_COMMIT_MESSAGE = $(M)

# T for Tag 
#GITR_TAG_VERSION ?= NULL
GITR_TAG_VERSION ?= $(T)




## Prints the git setting
gitr-print:
	@echo
	@echo -- GITR Upstream --
	@echo GITR_ORG_UPSTREAM: 			$(GITR_ORG_UPSTREAM)
	@echo GITR_REPO_UPSTREAM_ABS_URL: 	$(GITR_REPO_UPSTREAM_ABS_URL)
	
	@echo -- GITR Fork --
	@echo GITR_ORG_FORK: 				$(GITR_ORG_FORK)
	@echo GITR_SERVER: 					$(GITR_SERVER)
	@echo GITR_USER: 					$(GITR_USER)
	@echo GITR_REPO_NAME: 				$(GITR_REPO_NAME)

	@echo ---
	@echo GITR_REPO_ABS_URL: 			$(GITR_REPO_ABS_URL)
	@echo GITR_REPO_ABS_FSPATH: 		$(GITR_REPO_ABS_FSPATH)

	@echo ---
	@echo GITR_COMMIT_MESSAGE: 			$(GITR_COMMIT_MESSAGE)
	@echo GITR_TAG_VERSION: 			$(GITR_TAG_VERSION)
	@echo GITR_LAST_TAG:				$(GITR_LAST_TAG)
	
	@echo




### GIT-FORK

#See: https://help.github.com/en/github/collaborating-with-issues-and-pull-requests/syncing-a-fork

gitr-fork-all: gitr-status gitr-fork-catchup gitr-status gitr-fork-commit gitr-fork-push gitr-fork-open

gitr-status:
	# Test
	git status


### FORK

## Prints the exact Git Clone command you need :)
# You can call this from your Org folder:
# make -f shared/boilerplate/gitr.mk gitr-fork-clone-template
gitr-fork-clone-template:
	@echo
	@echo
	@echo Template is:
	@echo EX "git clone git@github.com-ME-getcouragenow:ME-getcouragenow/REPO_NAME"
	@echo
	@echo So if your fork is: 
	@echo github.com/james-getcouragenow/dev
	@echo
	@echo You use: 
	@echo EX: "git clone git@github.com-james-getcouragenow:james-getcouragenow/dev"
	@echo

## Sets up the git fork locally.
gitr-fork-setup-old:
	# Pre: you git forked ( via web) and git cloned (via ssh)
	# add upstream repo

	#git remote add upstream git://$(GITR_SERVER)/$(GITR_ORG_UPSTREAM)/$(GITR_REPO_NAME).git

## Sets up the git fork locally.
gitr-fork-setup:
	# Pre: you git forked ( via web) and git cloned (via ssh)
	# Sets up git config upstreak to point to the upstream origin
	@echo
	@echo EX git remote add upstream git@github.com-joe-getcouragenow:getcouragenow/dev
	@echo
	@echo EX git remote add upstream git@$(GITR_SERVER)-$(GITR_USER):$(GITR_ORG_UPSTREAM)/$(GITR_REPO_NAME)
	@echo
	# WORKS
	git remote add upstream git@$(GITR_SERVER)-$(GITR_USER):$(GITR_ORG_UPSTREAM)/$(GITR_REPO_NAME)
	@echo

## Sync upstream with your fork. Use this to make a PR.
gitr-fork-catchup:
	
	# This fetches the branches and their respective commits from the upstream repository.
	@echo
	git fetch upstream
	@echo

	# This brings your fork's master branch into sync with the upstream repository, without losing your local changes.
	@echo
	git merge upstream/$(GITR_BRANCH_NAME)
	@echo

## Commit the changes to the repo
gitr-fork-commit:
	@echo GITR_COMMIT_MESSAGE: $(GITR_COMMIT_MESSAGE)
	git add --all
	git commit -m '$(GITR_COMMIT_MESSAGE)'

## Push the repo to orgin
gitr-fork-push:
	git push origin $(GITR_BRANCH_NAME)

## Opens the forked git server.
gitr-fork-open:
	open $(GITR_REPO_ABS_URL).git

## Submits the PR you pushed
gitr-fork-pr-submit:
	## TODO. Alex gave me the commands.
	open $(GITR_REPO_ABS_URL).git

### UPSTREAM

## Opens the upstream git web
gitr-upstream-open:
	open https://$(GITR_SERVER)/$(GITR_ORG_UPSTREAM)/$(GITR_REPO_NAME).git 

gitr-upstream-merge:
	# Use github cli




## GIT-TAG

## Create a tag.
gitr-tag-create:
	# this will create a local tag on your current branch and push it to Github.

	git tag $(GITR_TAG_VERSION)

	# push it up
	#git push origin --tags

## Deletes a tag.
gitr-tag-delete:
	# this will delete a local tag and push that to Github

	# local
	git tag -d $(GITR_TAG_VERSION)

	# remote
	#git push --delete origin $(GITR_TAG_VERSION)

	

## GIT-RELEASE

## Stage a release (usage: make release-tag VERSION={VERSION_TAG})
gitr-release-tag:
	@echo Tagging release with version "${VERSION}"
	@git tag -a ${VERSION} -m "chore: release version '${VERSION}'"
	@echo Generating changelog
	@git-chglog -o CHANGELOG.md
	@git add CHANGELOG.md
	@git commit -m "chore: update changelog for version '${VERSION}'"

## Push a release (warning: ensure the release was staged first)
gitr-release-push: 
	@echo Publishing release
	@git push --follow-tags