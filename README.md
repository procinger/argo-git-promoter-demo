
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 15faabd5a73482b8355d54b8e6c6c027b7e33ad9
kustomize build ./kustomize/overlays/dev
```