
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout f845db29d31ded0af302f547221574ed895c23d7
kustomize build ./kustomize/overlays/test
```