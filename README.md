
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 3b986e215915beefde2bf625f9c3321ee85519ef
kustomize build ./kustomize/overlays/test
```