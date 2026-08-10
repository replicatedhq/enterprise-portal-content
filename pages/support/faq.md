---
title: Frequently Asked Questions
---

# Frequently Asked Questions

## General

<Accordion title="How do I check for updates?">

See [Instances & Updates](../updates/instances).

</Accordion>

{{#if entitlements.isHelmInstallEnabled}}
## Installation

<Accordion title="Which installation method should I use?">

{{#if entitlements.isEmbeddedClusterDownloadEnabled}}
Choose [Embedded Cluster (Linux)](../installation/linux) for installing on a Linux server, or [Helm](../installation/helm) for deploying to an existing Kubernetes cluster.
{{/if}}

</Accordion>
{{/if}}

## Troubleshooting

<Accordion title="How do I collect diagnostic information?">

See [Support Bundles](./bundles) to generate a new bundle, or [upload an existing one](./bundles#upload-an-existing-bundle).

</Accordion>
