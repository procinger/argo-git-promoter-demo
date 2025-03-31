
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout fd08c0ece8e38a1b97de9dab1ff79b2d39c41d90
kustomize build ./kustomize/overlays/test
```