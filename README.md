
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout fad89f567a94bba1d555b40045ee822e0eed0281
kustomize build ./kustomize/overlays/test
```