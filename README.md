
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout f673ad01696460644d7a802ef2c414f6963fe554
kustomize build ./kustomize/overlays/test
```