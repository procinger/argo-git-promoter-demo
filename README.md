
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout dc56202f0ababbd06a46a438e993540a6e4cea56
kustomize build ./kustomize/overlays/test
```