
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 32d74807c57a8bccf7419f9c5e52121aa0f608aa
kustomize build ./kustomize/overlays/test
```