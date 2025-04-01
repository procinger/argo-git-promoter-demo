
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 1a5652e4c6251fcd854102c13f9812ef557efda8
kustomize build ./kustomize/overlays/test
```