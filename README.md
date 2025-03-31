
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout cea3051adbcfa2d19c80df4c109db169d3fdd84a
kustomize build ./kustomize/overlays/dev
```