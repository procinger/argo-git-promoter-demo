
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout d89a53d6df1e82cc9bb09b40a7dd6537328cad93
kustomize build ./kustomize/overlays/dev
```