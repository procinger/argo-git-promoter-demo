
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 330eb3145e36ddd9997befef7385dc801413624f
kustomize build ./kustomize/overlays/test
```