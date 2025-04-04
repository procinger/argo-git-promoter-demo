
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 1ccef06b185eb446b2520bc7ae3b8e0403cf816b
kustomize build ./kustomize/overlays/test
```