
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout d56c69b6f3e57d984e4490258e871743a53b9245
kustomize build ./kustomize/overlays/dev
```