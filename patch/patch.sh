#!/bin/bash
set -ex

git config --global user.name "John Doe"
git config --global user.email "johndoe@example.com"

function git_cherry_pick ()
{
	refs=$1
	git fetch "https://gerrit.fd.io/r/vpp" ${refs}
	git cherry-pick FETCH_HEAD
	git commit --amend -m "gerrit:${refs#refs/changes/*/} $(git log -1 --pretty=%B)"
}

# --------------- Dedicated plugins ---------------

git_cherry_pick refs/changes/64/33264/7 # 33264: pbl: Port based balancer | https://gerrit.fd.io/r/c/vpp/+/33264

git_cherry_pick refs/changes/88/31588/4 # 31588: cnat: [WIP] no k8s maglev from pods | https://gerrit.fd.io/r/c/vpp/+/31588

git_cherry_pick refs/changes/83/28083/21 # 28083: acl: acl-plugin custom policies |  https://gerrit.fd.io/r/c/vpp/+/28083

if [ "$(ls ./patch/*.patch 2> /dev/null)" ]; then
  git apply patch/*.patch
  git add --all
  git commit -m "misc patches"
fi
