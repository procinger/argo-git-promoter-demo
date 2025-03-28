
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 8e689665d99fff0683f12265cd30e1fea93bdfbd
kustomize build ./kustomize/overlays/test
```