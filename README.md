
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout d268fb400f1de56354570d4822930cfc400a2d0c
kustomize build ./kustomize/overlays/test
```