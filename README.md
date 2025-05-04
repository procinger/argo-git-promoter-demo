
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 69cf11bd57f173e4f7f60e6de674370f9105a952
kustomize build ./kustomize/overlays/dev
```