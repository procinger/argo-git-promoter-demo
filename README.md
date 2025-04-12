
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout ac4bafe09c6b4dd5f83476a600a301afdc2e6bca
kustomize build ./kustomize/overlays/dev
```