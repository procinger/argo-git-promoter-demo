
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 0331a815dcde8600be192273a43b2b594ea98efd
kustomize build ./kustomize/overlays/dev
```