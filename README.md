
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout f6d622696bacaad9367edc5f9a1de0ea9783fcc1
kustomize build ./kustomize/overlays/dev
```