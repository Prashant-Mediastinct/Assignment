
if [ "$TRAVIS_PULL_REQUEST" == "false" ]; then 
    echo "Here"
    exit 0;
fi

if [ "$TRAVIS_BRANCH" == "develop" ]; then
echo "Merge to Develop"
git checkout "$TRAVIS_BRANCH" || exit 
git merge "$TRAVIS_COMMIT" || exit 
git push origin develop  # here need some authorization and url

echo "Merging Complete!";
fi
