
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout db43bd1e436238a7eee57b6a184cac9afe571f9c
kustomize build ./kustomize/overlays/dev
```