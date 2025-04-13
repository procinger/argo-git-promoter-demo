
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 730c2154dc0be6d7553baf433c75f720fd601839
kustomize build ./kustomize/overlays/test
```