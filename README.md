
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 22a5f097e8f18b43cd77a5eab1378c769931e41d
kustomize build ./kustomize/overlays/test
```