
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 4341b9d4398cbf817e1996dc61093a7ba894aaf6
kustomize build ./kustomize/overlays/test
```