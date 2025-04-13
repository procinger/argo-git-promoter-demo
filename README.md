
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout cf46652030a553e37f3c9a2cd36cd3e1701007d4
kustomize build ./kustomize/overlays/test
```