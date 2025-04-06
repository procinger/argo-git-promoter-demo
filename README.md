
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout e5b2510ade4b88a6247c4d2bd7384e7b7f5327f3
kustomize build ./kustomize/overlays/test
```