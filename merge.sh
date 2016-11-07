
if [ "$TRAVIS_PULL_REQUEST" == "false" ]; then 
    echo "Here"
    exit 0;
fi

if [ "$TRAVIS_BRANCH" == "develop" ]; then
echo "Merge to Develop"
git checkout "$TRAVIS_BRANCH" || exit 1
git merge "$TRAVIS_COMMIT" || exit 1
git push origin feature  # here need some authorization and url

echo "Merging Complete!";
fi
