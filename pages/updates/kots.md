---
title: Upgrade KOTS Applications
visible_when:
  entitlements:
    - isKotsInstallEnabled
---

# Upgrade KOTS Applications

Use the Admin Console to upgrade an application installed with KOTS. Enterprise
Portal helps you choose a target release and, for air gap instances, download
the application bundle. The deployment itself happens in Admin Console, and
Enterprise Portal does not track its progress or completion.

> **Vendor customization:** Replace the generic Admin Console access, private
> registry, and support instructions below with the details for your
> environment. You can also change or remove this page and its link from the
> Instances & Updates panel.

## Choose the target release

1. Return to [Instances & Updates](/updates/instances).
2. Select the KOTS instance you want to upgrade.
3. Review the available releases and select the intended target. Confirm in
   Admin Console that the release is available before starting the deployment.

## Upgrade an online or proxy instance

1. Reopen the KOTS Admin Console by running
   `kubectl kots admin-console --namespace {{app.slug}}`.
2. Open the **Version History** tab. The Admin Console checks for new versions
   every four hours by default, so the target release may already be listed.
   If it is not listed, click **Check for updates**.
3. Select the same target release you chose in Enterprise Portal. Review the
   release notes, configuration changes, and preflight checks.
4. Click **Deploy** when you are ready to upgrade the application.
5. Monitor the deployment in Admin Console. Contact your vendor's support team
   if the release is unavailable or a check fails.

`kubectl kots upstream upgrade` checks for updates, and the customer still
deploys in the Admin Console. Document that command here only when every
customer uses the same namespace. Remove it otherwise.

```
kubectl kots upstream upgrade --namespace {{app.slug}}
```

{{#if entitlements.isAirgapSupported}}

## Upgrade an air gap instance

1. In [Instances & Updates](/updates/instances), select the target release and
   download its **Application air gap bundle**. Move the file into the
   air-gapped environment using your approved transfer process.
2. If you need to reopen the Admin Console, run `kubectl kots admin-console --namespace {{app.slug}}`.
   Then open **Version History**.
3. Upload the application `.airgap` bundle. If your environment uses a private
   registry, follow your vendor's registry-specific preparation instructions.
4. Select the uploaded release, review configuration changes and preflight
   checks, and click **Deploy**.
5. Monitor the deployment in Admin Console. Enterprise Portal does not receive
   completion status from the instance.

{{/if}}

## Get help

If you cannot access Admin Console, do not see the target release, or cannot
complete a required check, use the support process supplied by your vendor.
