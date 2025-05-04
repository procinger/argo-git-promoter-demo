
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 52321c7692aac9a9a877dacb4d510c716b0cbb15
kustomize build ./kustomize/overlays/dev
```