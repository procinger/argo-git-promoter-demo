
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout ad813a7a17a4e4089247bc6d4843283b14b9bbb7
kustomize build ./kustomize/overlays/test
```