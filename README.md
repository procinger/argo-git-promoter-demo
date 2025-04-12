
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 68952b55d3aa73f536765413e574bdec5b3ed822
kustomize build ./kustomize/overlays/dev
```