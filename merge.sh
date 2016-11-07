
if [ "$TRAVIS_PULL_REQUEST" == "false" ]; then 
    echo "Here"
    exit 0;
fi

if [ "$TRAVIS_BRANCH" == "develop" ]; then
echo "Merge to Develop"
<<<<<<< HEAD
git checkout "$TRAVIS_BRANCH" || exit 1
git merge "$TRAVIS_COMMIT" || exit 1
=======
git checkout "$TRAVIS_BRANCH" || exit 
git merge "$TRAVIS_COMMIT" || exit 
>>>>>>> 100712af7d935a8dc67605a6df5b0fc706b20989
git push origin feature  # here need some authorization and url

echo "Merging Complete!";
fi
