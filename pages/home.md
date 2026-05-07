---
title: Documentation
---

# Documentation

Welcome to your personalized documentation portal. The navigation and content are customized based on your license entitlements.

## Available Features

Your installation includes access to the following features:

{{#if entitlements.isEmbeddedClusterDownloadEnabled}}
- **Linux (Embedded Cluster):** Install on a Linux server using Embedded Cluster
{{/if}}
{{#if entitlements.isHelmInstallEnabled}}
- **Helm Installation:** Deploy to existing Kubernetes clusters using Helm charts
{{/if}}
{{#if entitlements.isAirgapSupported}}
- **Air Gap Support:** Install in disconnected environments
{{/if}}

## Getting Started

<Tip title="New to this portal?">
Start with the Installation Guide for your deployment method. The configuration selectors on each installation page will generate customized commands for your environment.
</Tip>

Use the sidebar navigation on the left to explore available documentation sections. We recommend starting with:

{{#if entitlements.isEmbeddedClusterDownloadEnabled}}
1. **Installation Requirements** — Review system requirements and prerequisites for Embedded Cluster installations
{{/if}}
2. **Installation Guide** — Follow step-by-step installation instructions for your deployment method
3. **Updates** — Check for and manage application updates
4. **Support Bundles** — Generate diagnostic bundles for troubleshooting

## Quick Links

<OptionSelector label="Install Method" defaultOption="Linux" storageKey="install-method">
<Option value="Linux">

{{#if entitlements.isEmbeddedClusterDownloadEnabled}}
- [Installation Requirements](installation/requirements)
- [Linux Installation](installation/linux)
{{/if}}
- [Release History](installation/release-history)
- [Instances & Updates](updates/instances)
- [Support Bundles](support/bundles)
- [FAQ](support/faq)

</Option>
<Option value="Helm">

{{#if entitlements.isHelmInstallEnabled}}
- [Helm Installation](installation/helm)
{{/if}}
- [Release History](installation/release-history)
- [Instances & Updates](updates/instances)
- [Support Bundles](support/bundles)
- [FAQ](support/faq)

</Option>
</OptionSelector>
