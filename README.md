
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 99aa22de1057dfcca10a4ffba783b9d381b3df4d
kustomize build ./kustomize/overlays/dev
```