
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 20d5242452c8d70062606033592604958a7bd1c5
kustomize build ./kustomize/overlays/dev
```