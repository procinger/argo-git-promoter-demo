
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 73fbde0320d1ba3a3cdc5891a36c669292844936
kustomize build ./kustomize/overlays/test
```