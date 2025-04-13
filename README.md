
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout da0aeace3b70fa80ca84ff75c2f6ca595a04f059
kustomize build ./kustomize/overlays/test
```