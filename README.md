
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout f496165ced73785483e33db0763f025e11acfbb6
kustomize build ./kustomize/overlays/test
```