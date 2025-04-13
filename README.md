
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 278999fd8a3b6a5a3b1bb4576da32c246b514be5
kustomize build ./kustomize/overlays/dev
```