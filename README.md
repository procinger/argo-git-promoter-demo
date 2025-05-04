
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout ec9610bd5daed60c595fcd987f501a74fbf4a072
kustomize build ./kustomize/overlays/dev
```