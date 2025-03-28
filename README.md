
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 585a2b5c4c4e6be41784880daefef4e8849a444c
kustomize build ./kustomize/overlays/dev
```