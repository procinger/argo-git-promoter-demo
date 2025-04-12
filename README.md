
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout f75386d67632a3015219d8e21cdf271feea87947
kustomize build ./kustomize/overlays/test
```