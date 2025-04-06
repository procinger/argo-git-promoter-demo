
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout ee158b2054a31dad9c4b547ea1e455f26817747e
kustomize build ./kustomize/overlays/test
```