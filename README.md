
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 339db718a511f1f05cc7135bdf08a55e7f360344
kustomize build ./kustomize/overlays/test
```