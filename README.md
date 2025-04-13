
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 0436c935975a1ae476ec91fa1b6b95dbc3eb4d14
kustomize build ./kustomize/overlays/test
```