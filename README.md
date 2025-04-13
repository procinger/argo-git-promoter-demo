
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout ea489cd5e174e7cc9da5457312a25fcb9850c7e0
kustomize build ./kustomize/overlays/dev
```