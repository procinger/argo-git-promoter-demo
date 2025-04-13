
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 1ff643c049a9f6a85409ca5d6c55112077761746
kustomize build ./kustomize/overlays/dev
```