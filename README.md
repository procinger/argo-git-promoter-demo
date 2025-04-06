
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 0a14575b7ae3ad5511140c9b8aba7a4f46c68583
kustomize build ./kustomize/overlays/test
```