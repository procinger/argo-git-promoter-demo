
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 05f2c68ca90142eb7874611910aded5e9dd97f9b
kustomize build ./kustomize/overlays/dev
```