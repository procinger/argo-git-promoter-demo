
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 9ecfb91effc87d40b2861a7468f1d7520c262552
kustomize build ./kustomize/overlays/test
```