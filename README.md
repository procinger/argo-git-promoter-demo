
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout ac6a3fb9d184753e98ddbd95e9f86a739473a626
kustomize build ./kustomize/overlays/test
```