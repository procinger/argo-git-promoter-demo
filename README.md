
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 8e90b2cdabb994ce9e75aae78d8eaea9d07f8aff
kustomize build ./kustomize/overlays/dev
```