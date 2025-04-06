
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout f61f7a2ac0210bdb4bf9840204007e986c0717e6
kustomize build ./kustomize/overlays/test
```