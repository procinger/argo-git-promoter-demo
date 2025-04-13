
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 88d365a779efea6825a9b7435344e2d393035af0
kustomize build ./kustomize/overlays/test
```