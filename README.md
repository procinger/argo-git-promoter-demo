
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 63b8233ac2efaeab402e9b9961081d627012ab91
kustomize build ./kustomize/overlays/test
```