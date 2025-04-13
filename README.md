
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 4d1d02a019624d057040e550c3574f4fa0889845
kustomize build ./kustomize/overlays/test
```