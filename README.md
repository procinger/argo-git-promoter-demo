
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 6fc36cf10eb4f06cca2dd276defcbcfc5ec14861
kustomize build ./kustomize/overlays/test
```