
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 0ee2b844d19a3ae66803e46fe33febd2538fba0f
kustomize build ./kustomize/overlays/dev
```