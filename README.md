
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 28abc10fa02b6530dbcec8e816143803987f3914
kustomize build ./kustomize/overlays/test
```