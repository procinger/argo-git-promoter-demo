
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 41fa8804b33a416ab694780ca918a25775eac702
kustomize build ./kustomize/overlays/test
```