
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 515fae771533a4f47d121d5da2425126a7d9c876
kustomize build ./kustomize/overlays/dev
```