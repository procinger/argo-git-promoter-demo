
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 609aadf979318d07642eeb48dd5e7b152b2c629d
kustomize build ./kustomize/overlays/test
```