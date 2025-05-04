
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 0940691ac09745703f84f16a449a77d68c7fcaeb
kustomize build ./kustomize/overlays/dev
```