
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 4031064d0245f7fce7a0f8d0f5ea830fb97db997
kustomize build ./kustomize/overlays/dev
```