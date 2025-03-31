
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 3953aa6bbcbd9330c8d98e87baddfb44703f6611
kustomize build ./kustomize/overlays/test
```