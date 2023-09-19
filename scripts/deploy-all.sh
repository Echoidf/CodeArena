#!/bin/bash

#获取当前分支名称
cur_branch=$(git symbolic-ref --short HEAD)

the_other_branch=""
if [ $cur_branch = "main" ]
then
  the_other_branch="business"
else
  the_other_branch="main"
fi

#提交代码到远程对应分支
git push origin $cur_branch
#提交代码到远程另一个分支
git push origin $the_other_branch

#切换到另一个本地分支
git checkout $the_other_branch
#更新代码
git cherry-pick -X theirs $commit_hash
git push origin $the_other_branch

git checkout $cur_branch