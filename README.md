
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 9a15c36808c5b26d45ebb72818581248ac949df7
kustomize build ./kustomize/overlays/dev
```