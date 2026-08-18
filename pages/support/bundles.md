---
title: Support Bundles
---

# Support Bundles

<Tip>
Generate a support bundle whenever you encounter unexpected behavior. Bundles capture system state, logs, and configuration at a point in time, which is critical for diagnosing intermittent issues.
</Tip>

## Generate a Bundle

<Tabs>
{{#if entitlements.isEmbeddedClusterDownloadEnabled}}
<Tab title="Linux (Embedded Cluster)">

<LinuxBundles />

</Tab>
{{/if}}
{{#if entitlements.isKurlInstallEnabled}}
<Tab title="Linux (kURL)">

<KurlBundles />

</Tab>
{{/if}}
{{#if entitlements.isHelmInstallEnabled}}
<Tab title="Existing Cluster (Helm)">

<HelmBundles />

</Tab>
{{/if}}
{{#if entitlements.isKotsInstallEnabled}}
<Tab title="Existing Cluster (KOTS)">

<KotsBundles />

</Tab>
{{/if}}
</Tabs>

## Upload an Existing Bundle

<SupportBundleUpload />

## Uploaded Bundles

<SupportBundleUploadHistory />

<ContactInfo />
