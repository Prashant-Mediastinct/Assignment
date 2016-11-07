
if [ "$TRAVIS_PULL_REQUEST" == "false" ] || [ "$TRAVIS_BRANCH" != "feature" ]; then 
    exit 0;
fi


git checkout develop || exit
git merge "$TRAVIS_COMMIT" || exit
git push origin "$TRAVIS_PULL_REQUEST_BRANCH"  # here need some authorization and url

