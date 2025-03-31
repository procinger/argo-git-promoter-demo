
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 9f38112518087222297f2193bc78076c1dfb6308
kustomize build ./kustomize/overlays/dev
```