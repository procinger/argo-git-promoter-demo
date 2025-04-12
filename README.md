
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 22234757a6f8b5886c1d067dcd1403733ff89d43
kustomize build ./kustomize/overlays/dev
```