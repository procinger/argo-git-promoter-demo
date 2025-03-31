
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 0a2e9bb626a656f2c91746ef2a535ff3207bf3a7
kustomize build ./kustomize/overlays/dev
```