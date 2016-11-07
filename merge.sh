
if [ "$TRAVIS_PULL_REQUEST" == "false" ]; then 
    echo "Here"
    exit 0;
fi

echo "Time to merge"
git checkout develop || exit
git merge "$TRAVIS_COMMIT" || exit
git push origin feature  # here need some authorization and url

echo "Merging Complete!"

#  || [ "$TRAVIS_BRANCH" != "feature" ]