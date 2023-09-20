#!/bin/bash

#获取当前分支名称
cur_branch=$(git symbolic-ref --short HEAD)
echo "====== 当前分支: " $cur_branch "======"

the_other_branch=""
if [ $cur_branch = "main" ]
then
  the_other_branch="business"
else
  the_other_branch="main"
fi

#提交代码到远程对应分支
echo $cur_branch" 代码提交中..."
git push origin $cur_branch

commit_hash=$(git log -n 1 --pretty=format:"%H")

#切换到另一个本地分支
git checkout $the_other_branch
echo "====== 当前分支:" $the_other_branch "======"
#更新代码
git cherry-pick -X theirs $commit_hash
echo $cur_branch" 代码提交中..."
git push origin $the_other_branch

git checkout $cur_branch