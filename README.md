
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout b1109ff6e418e1aa41f8f01504ec4994cc5c7630
kustomize build ./kustomize/overlays/dev
```