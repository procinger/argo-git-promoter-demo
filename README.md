
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 8e1c5c63aa33e4ac0b16320084dbbbd564f2167c
kustomize build ./kustomize/overlays/dev
```