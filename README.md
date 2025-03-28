
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout bf6907cbab848ba104a7e09b35f03bd80ba2a9c9
kustomize build ./kustomize/overlays/test
```