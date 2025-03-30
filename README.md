
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout f0fac7f934c49266c78b0bde11c61a68b4e9c13f
kustomize build ./kustomize/overlays/test
```