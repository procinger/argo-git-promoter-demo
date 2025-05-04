
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 6662bc19c64881d455dac63b26815bfc46d76b26
kustomize build ./kustomize/overlays/test
```