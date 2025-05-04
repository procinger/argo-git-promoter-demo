
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout ae40744aa7114bb906068f18abb581af78b5489f
kustomize build ./kustomize/overlays/dev
```