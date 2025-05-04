
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 565b1807ad0acd695290496da3106fce1daaa556
kustomize build ./kustomize/overlays/dev
```