
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 6b311dce997620096c17b1f845398822267d2a31
kustomize build ./kustomize/overlays/test
```