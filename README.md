
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout f0cb0cdc480ebe31da57e9d8dddf79c8974f27f8
kustomize build ./kustomize/overlays/dev
```