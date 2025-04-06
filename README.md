
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout bd8c937efd7e26845ee85e759a0dcbf4c7a2adb9
kustomize build ./kustomize/overlays/test
```