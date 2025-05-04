
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout f83c9a30825c9ded59a18fb5ade746a4e94771a6
kustomize build ./kustomize/overlays/test
```