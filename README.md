
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 1ed3de52dd570484b1cabb6ab07729c9cb305484
kustomize build ./kustomize/overlays/dev
```