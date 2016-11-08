
if [ "$TRAVIS_PULL_REQUEST" == "false" ]; then 
    echo "Here"
    exit 0;
fi

export GIT_COMMITTER_EMAIL="prashant@mediastinct.com"
export GIT_COMMITTER_NAME="Prashant-Mediastinct"

if [ "$TRAVIS_BRANCH" == "develop" ]; then
echo "Merge to Develop"
# git config user.email "prashant@mediastinct.com"
# git config user.name "Prashant-Mediastinct"
git fetch --all --prune
echo "after pull"
git checkout "origin/$TRAVIS_BRANCH" || exit 
echo "after checkout"
git merge "origin/$TRAVIS_COMMIT" || exit 
echo "after merge"
git push origin "origin/$TRAVIS_BRANCH"  # here need some authorization and url

echo "Merging Complete!";
fi
