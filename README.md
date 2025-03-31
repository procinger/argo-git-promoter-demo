
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 8eb1e11fa66e33d61e61a13403043a9302a8ecda
kustomize build ./kustomize/overlays/dev
```