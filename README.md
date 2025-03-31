
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout c039d4b272583841ab6a50e347d880a4dcab8605
kustomize build ./kustomize/overlays/dev
```