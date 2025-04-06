
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout b99b2d0bcfe50ee949ebd92f3fa078d63c5388ad
kustomize build ./kustomize/overlays/test
```