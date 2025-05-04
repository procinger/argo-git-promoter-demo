
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 54b188ff6e7e045d62b45f176f09ca40fe8dbf5b
kustomize build ./kustomize/overlays/test
```