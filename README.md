
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 3c731204707919be0ec2a4fddb6287f6d6d64a75
kustomize build ./kustomize/overlays/test
```