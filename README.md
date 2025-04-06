
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 473b4da0b908d362d6023a6675f978e671a5897b
kustomize build ./kustomize/overlays/test
```