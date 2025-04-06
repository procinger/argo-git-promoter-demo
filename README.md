
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout e837f999138aca7c65a73e111c2916f8b30e026b
kustomize build ./kustomize/overlays/dev
```