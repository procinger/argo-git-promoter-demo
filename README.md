
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 912987c584c5abe449696d24a5446cf65d2a5382
kustomize build ./kustomize/overlays/dev
```