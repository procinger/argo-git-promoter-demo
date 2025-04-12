
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 0e4df14f6049dc19633cadba068a9a55be19d84c
kustomize build ./kustomize/overlays/test
```