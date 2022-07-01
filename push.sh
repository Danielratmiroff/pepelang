#!/bin/bash

# Script to commit & push current branch with a default message if none is provided
# @Param #1: commit message

BRANCH=$(git rev-parse --abbrev-ref HEAD)

COMMENT="$1"
COMMIT_COMMENT="${COMMENT:-"Progress"}"

function push {
    echo "
----------------------------------------------
Running tests
----------------------------------------------
	"

    go test ./...

    if [ "$?" == "1" ]; then
        while true; do
            read -p "
----------------------------------------------
Tests FAILED.
Do you wish to continue pushing to master?
----------------------------------------------
	" yn
            case $yn in
            [Yy]*)
                echo "Wild decission!, but ok"
                break
                ;;
            [Nn]*)
                echo "Ok, exiting now."
                exit
                ;;
            *) echo "Please answer yes or no." ;;
            esac
        done
    fi

    echo "
----------------------------------------------
Committing changes: "${COMMIT_COMMENT}"
----------------------------------------------
	"

    git add .
    git commit -m "${COMMIT_COMMENT}"

    echo "
----------------------------------------------
Pushing to "${BRANCH}", with Commit: "${COMMIT_COMMENT}"
----------------------------------------------
	"
    git push origin ${BRANCH}
    exit
}

if [ "${BRANCH_PARAM}" = "--help" ]; then
    echo "push.sh \"Commit-Message\""
    exit 1
else
    push
fi
