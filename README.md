
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout d97736d025f0a76d32a00d9b1731ff8fe66a88fb
kustomize build ./kustomize/overlays/test
```