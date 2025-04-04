
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout f5ab3e61cc9f87a14e0482de6f9179cffbc96642
kustomize build ./kustomize/overlays/dev
```