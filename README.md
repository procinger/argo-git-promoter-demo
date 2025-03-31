
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout d4abb44198fed203a7b5924336e40755e0a10058
kustomize build ./kustomize/overlays/test
```