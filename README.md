
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 0600f3b0755921000b211a5c6e18aa66aff5289a
kustomize build ./kustomize/overlays/test
```