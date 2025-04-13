
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout a2e46cfeec6ce6a72eb23201218b6a0583290339
kustomize build ./kustomize/overlays/test
```