
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 25eb89483207db70e857787bb06675849d05f0c4
kustomize build ./kustomize/overlays/test
```