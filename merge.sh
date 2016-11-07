
if [ "$TRAVIS_PULL_REQUEST" == "false" ] || [ "$TRAVIS_BRANCH" != "feature" ]; then 
    exit 0;
fi

export GIT_COMMITTER_EMAIL=...
export GIT_COMMITTER_NAME=...

git checkout develop || exit
git merge "$TRAVIS_COMMIT" || exit
git push origin "$TRAVIS_PULL_REQUEST_BRANCH"  # here need some authorization and url

