
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout a2bcd0c39c6c4a5fbebb9ccbd2e6e56a868de486
kustomize build ./kustomize/overlays/test
```