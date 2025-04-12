
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout fd7228f60f1ebf3a3384b7ad2caadfdd4c18bfd4
kustomize build ./kustomize/overlays/dev
```