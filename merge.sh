
if [ "$TRAVIS_PULL_REQUEST" == "false" ]; then 
    echo "Here"
    exit 0;
fi

if [ "$TRAVIS_BRANCH" == "develop" ]; then
echo "Merge to Develop"
# git config user.email "prashant@mediastinct.com"
# git config user.name "Prashant-Mediastinct"
git pull origin
echo "after pull"
git checkout "$TRAVIS_BRANCH" || exit 
echo "after checkout"
git merge "$TRAVIS_COMMIT" || exit 
echo "after merge"
git push origin "$TRAVIS_BRANCH"  # here need some authorization and url

echo "Merging Complete!";
fi
