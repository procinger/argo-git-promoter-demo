
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout e20ac6e655d0954916bb59e947b2f127fdadeef6
kustomize build ./kustomize/overlays/test
```