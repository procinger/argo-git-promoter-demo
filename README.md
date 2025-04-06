
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout f8e2fd466e9c780e43f560eb6b550b88bebd803b
kustomize build ./kustomize/overlays/test
```