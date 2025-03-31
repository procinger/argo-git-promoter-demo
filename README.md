
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 569d6264378d20baed18101f6d629dbf365cd605
kustomize build ./kustomize/overlays/test
```