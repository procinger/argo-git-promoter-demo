
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout dde235aee108f109f84db8e07714bc78ab1e1eff
kustomize build ./kustomize/overlays/test
```