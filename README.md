
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 42f4885cbc05112c5dec01cc1d6e7433c0e0d315
kustomize build ./kustomize/overlays/test
```