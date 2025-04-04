
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout f7f67c7b8c1754e358087e8f86d56f348359769d
kustomize build ./kustomize/overlays/dev
```