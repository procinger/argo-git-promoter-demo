
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout bd0b90d9ab801bcebcf971ff14e8fb7e082a850b
kustomize build ./kustomize/overlays/dev
```