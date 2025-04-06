
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 8e2fffd482438d2ba11b3a9a1e39fad4a70b0e57
kustomize build ./kustomize/overlays/dev
```