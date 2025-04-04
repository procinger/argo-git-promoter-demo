
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 3b189797489ba094cb90d96269474b6162e0e78f
kustomize build ./kustomize/overlays/dev
```