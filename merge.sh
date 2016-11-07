
if [ "$TRAVIS_PULL_REQUEST" == "false" ]; then 
    echo "Here"
    exit 0;
fi

if [ "$TRAVIS_BRANCH" == "develop" ]; then
echo "Merge to Develop"
git clone git@github.com:Prashant-Mediastinct/Assignment.git
echo "after clone"
git checkout "$TRAVIS_BRANCH" || exit 
echo "after checkout"
git merge "$TRAVIS_COMMIT" || exit 
echo "after merge"
git push origin "$TRAVIS_BRANCH"  # here need some authorization and url

echo "Merging Complete!";
fi