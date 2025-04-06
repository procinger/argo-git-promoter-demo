
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 1f60e11d9ce01db317595a3c0f2aacd0ee5a6048
kustomize build ./kustomize/overlays/test
```