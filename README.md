
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout ed44e2e0ef6a011023d24441dc36a60e2bade853
kustomize build ./kustomize/overlays/dev
```