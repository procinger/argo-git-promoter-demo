
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 2a2da0cb0922986cb5e764e4457851ac2e9fd637
kustomize build ./kustomize/overlays/dev
```