#!/bin/bash
. ~/.nvm/nvm.sh
. ~/.profile
. ~/.bashrc

node="v18.7.0"
remote="origin"

origin="main"

branch="master"
# branch=bugfix/shipping-flow-bug-06
# branch=feature/sign-up-api-integrated
# branch=release/shell_script

prCommentResolverCommit=2

title="worked on otp verify & random otp generator, made constants for error messages and removed gitignore file"
description=""
developer="Abhinay Narayan Singh"
currentdate=`date +%d-%b-%Y`

if [$description == ""]
then
description=$title
fi


userConfirm=y
currentNodeVersion=$(node -v)
current_branch=$(git rev-parse --abbrev-ref HEAD)



# if [ $currentNodeVersion == $node ]
# then
# echo $currentNodeVersion
# else
# nvm use $node
# fi


# if [ $origin != $current_branch ]
# then
#     git checkout $origin
# fi
# git pull



exists=`git show-ref refs/heads/$branch`
if [ -n "$exists" ] 
then
    echo "$branch allready exists! wish to continue: y/n"
    read userConfirm
else
    git branch $branch
fi




if [ $userConfirm == "y" ] 
then
    git checkout $branch
    git add .

    if [ $prCommentResolverCommit == 1 ]
    then
        git commit -m "Title: Pull Request comment resolve; Developer: $developer; Date: $currentdate;"
    elif [ $prCommentResolverCommit == 2 ]
    then
        git commit -m "Title: $title; Developer: $developer; Date: $currentdate;"
    else
        git commit -m "Title: $title; Developer: $developer; Date: $currentdate;"
    fi
    
    git push --set-upstream $remote $branch

else
    echo "script execution terminated ..."
fi