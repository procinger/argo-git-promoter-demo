
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 362a119d846d7e4e92cfe82b92a25bcbffd718b0
kustomize build ./kustomize/overlays/test
```