
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout bfa17b6fc3edc8b57bb7500ec8e0df8410988d85
kustomize build ./kustomize/overlays/test
```