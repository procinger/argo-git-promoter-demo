
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 6a8db181442f1b5255d8c4e41f472d09b19ff09c
kustomize build ./kustomize/overlays/test
```