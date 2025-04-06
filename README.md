
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 52353025ec3771b3579a0314130ffbd91cc00d3a
kustomize build ./kustomize/overlays/test
```