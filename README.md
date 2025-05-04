
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 37e6706d086debe5aa2aa95a8dcbdbdcfc81c570
kustomize build ./kustomize/overlays/test
```